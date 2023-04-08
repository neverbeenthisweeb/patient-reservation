package app

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/rs/zerolog/log"
)

var (
	errNotFound                 = errors.New("not found")
	errReservationAlreadyExists = errors.New("reservation already exists")
)

type usecase interface {
	GetSlots(ctx context.Context) ([]slot, error)
	CreateReservation(ctx context.Context, patientID, doctorID, slotID int) (reservation, error)
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
		logger.Error().Err(errNotFound).Msg("Patient is not found")
		return reservation{}, errNotFound
	}

	// Does doctor exist?
	dct, err := uc.repo.doctor.GetDoctor(ctx, doctorID)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get doctor")
		return reservation{}, err
	}
	if reflect.DeepEqual(dct, doctor{}) {
		logger.Error().Err(errNotFound).Msg("Doctor is not found")
		return reservation{}, errNotFound
	}

	// Does doctor already have reservation?
	rv, err := uc.repo.reservation.GetReservations(
		ctx,
		getReservationsFilter{doctorID: doctorID},
	)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get reservations")
		return reservation{}, err
	}
	if len(rv) > 0 {
		logger.Error().Err(errReservationAlreadyExists).Msg("Reservation already exists")
		return reservation{}, errReservationAlreadyExists
	}

	// Does slot exist?
	sl, err := uc.repo.slot.GetSlot(ctx, slotID)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get slot")
		return reservation{}, err
	}
	if reflect.DeepEqual(sl, slot{}) {
		logger.Error().Err(errNotFound).Msg("Slot is not found")
		return reservation{}, errNotFound
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
			ID:        qn,
			PatientID: pt.ID,
			DoctorID:  dct.ID,
			// StartedAt: ,
			// EndedAt: ,
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

// func (uc *usecaseImpl) CancelReservation(ctx context.Context, reservationID string) error

// func (uc *usecaseImpl) GetReservations(ctx context.Context, showCancelled bool) ([]reservation, error)
