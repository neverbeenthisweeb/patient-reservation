package app

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/rs/zerolog/log"
)

var (
	errPatientNotFound      = errors.New("patient not found")
	errDoctorNotFound       = errors.New("doctor not found")
	errReservationNotFound  = errors.New("reservation not found")
	errSlotNotFound         = errors.New("slot not found")
	errReservationExists    = errors.New("reservation already exists")
	errReservationCancelled = errors.New("reservation is already cancelled")
)

type usecase interface {
	GetSlots(ctx context.Context) ([]slot, error)
	CreateReservation(ctx context.Context, patientID, doctorID, slotID int) (reservation, error)
	CancelReservation(ctx context.Context, reservationID int) (reservation, error)
	GetReservations(ctx context.Context, showCancelled bool) ([]reservation, error)
}

type usecaseImpl struct {
	repo  *repo
	infra *infra
}

func NewUsecase(repo *repo, infra *infra) *usecaseImpl {
	return &usecaseImpl{
		repo:  repo,
		infra: infra,
	}
}

func (uc *usecaseImpl) GetSlots(ctx context.Context) ([]slot, error) {
	logger := log.With().Str("requestid", ctx.Value("requestid").(string)).Logger()

	slots, err := uc.repo.slot.GetSlots(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get slots")
		return []slot{}, err
	}

	return slots, nil
}

func (uc *usecaseImpl) CreateReservation(ctx context.Context, patientID, doctorID, slotID int) (reservation, error) {
	logger := log.With().Str("requestid", ctx.Value("requestid").(string)).Logger()

	// Does patient exist?
	pt, err := uc.repo.patient.GetPatient(ctx, patientID)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get patient")
		return reservation{}, err
	}
	if reflect.DeepEqual(pt, patient{}) {
		logger.Error().Err(errPatientNotFound).Msg("Patient is not found")
		return reservation{}, errPatientNotFound
	}

	// Does doctor exist?
	dct, err := uc.repo.doctor.GetDoctor(ctx, doctorID)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get doctor")
		return reservation{}, err
	}
	if reflect.DeepEqual(dct, doctor{}) {
		logger.Error().Err(errDoctorNotFound).Msg("Doctor is not found")
		return reservation{}, errDoctorNotFound
	}

	// Does slot exist?
	sl, err := uc.repo.slot.GetSlot(ctx, slotID)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get slot")
		return reservation{}, err
	}
	if reflect.DeepEqual(sl, slot{}) {
		logger.Error().Err(errSlotNotFound).Msg("Slot is not found")
		return reservation{}, errSlotNotFound
	}

	// Does doctor already have reservation?
	rvs, err := uc.repo.reservation.GetReservations(
		ctx,
		getReservationsFilter{
			doctorID: doctorID,
			start:    sl.StartedAt,
			end:      sl.EndedAt,
		},
	)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get reservations")
		return reservation{}, err
	}
	if len(rvs) > 0 {
		logger.Error().Err(errReservationExists).Msg("Reservation already exists")
		return reservation{}, errReservationExists
	}

	// Get queue number
	qn, err := uc.infra.queueCounter.Count(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to count queue")
		return reservation{}, err
	}

	// Create reservation
	newRv, err := uc.repo.reservation.CreateReservation(
		ctx,
		reservation{
			ID:          qn,
			PatientID:   pt.ID,
			DoctorID:    dct.ID,
			StartedAt:   newSlotTime(sl.StartedAt),
			EndedAt:     newSlotTime(sl.EndedAt),
			IsCancelled: false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to create reservation")
		return reservation{}, err
	}

	return newRv, nil
}

func (uc *usecaseImpl) CancelReservation(ctx context.Context, ID int) (reservation, error) {
	logger := log.With().Str("requestid", ctx.Value("requestid").(string)).Logger()

	rv, err := uc.repo.reservation.GetReservation(ctx, ID)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get reservation")
		return reservation{}, err
	}
	if reflect.DeepEqual(rv, reservation{}) {
		logger.Error().Err(errReservationNotFound).Msg("Reservation is not found")
		return reservation{}, errReservationNotFound
	}
	if rv.IsCancelled {
		logger.Error().Err(errReservationCancelled).Msg("Reservation is canceled")
		return reservation{}, errReservationCancelled
	}

	cancelled, err := uc.repo.reservation.CancelReservation(ctx, ID)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to cancel reservation")
		return reservation{}, err
	}

	return cancelled, nil
}

func (uc *usecaseImpl) GetReservations(ctx context.Context, showCancelled bool) ([]reservation, error) {
	logger := log.With().Str("requestid", ctx.Value("requestid").(string)).Logger()

	rvs, err := uc.repo.reservation.GetReservations(
		ctx,
		getReservationsFilter{
			showCancelled: showCancelled,
		},
	)
	if err != nil {
		logger.Error().Err(err).Msg("Failed get reservations")
		return []reservation{}, err
	}

	return rvs, nil
}
