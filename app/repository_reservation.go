package app

import (
	"context"
	"time"
)

var (
	dbReservations = []reservation{
		{
			ID:          1,
			PatientID:   1,
			DoctorID:    1,
			StartedAt:   time.Date(baseTime.Year(), baseTime.Month(), baseTime.Day(), 10, 15, 0, 0, baseTime.Location()),
			EndedAt:     time.Date(baseTime.Year(), baseTime.Month(), baseTime.Day(), 10, 45, 0, 0, baseTime.Location()),
			IsCancelled: false,
			CreatedAt:   baseTime,
			UpdatedAt:   baseTime,
		},
		{
			ID:          2,
			PatientID:   2,
			DoctorID:    1,
			StartedAt:   time.Date(baseTime.Year(), baseTime.Month(), baseTime.Day(), 11, 30, 0, 0, baseTime.Location()),
			EndedAt:     time.Date(baseTime.Year(), baseTime.Month(), baseTime.Day(), 13, 30, 0, 0, baseTime.Location()),
			IsCancelled: true,
			CreatedAt:   baseTime,
			UpdatedAt:   baseTime,
		},
	}
)

type repoReservationImpl struct{}

func NewRepoReservationImpl() *repoReservationImpl {
	return &repoReservationImpl{}
}

func (r *repoReservationImpl) GetReservations(ctx context.Context, filter getReservationsFilter) ([]reservation, error) {
	ret := []reservation{}

	for _, v := range dbReservations {
		ok := true

		if filter.doctorID != 0 && v.DoctorID != filter.doctorID {
			ok = false
		}

		if !filter.showCancelled && v.IsCancelled {
			ok = false
		}

		if (filter.start != "" && !isWithinTimeSlot(filter.start, v.StartedAt, v.EndedAt)) &&
			(filter.end != "" && !isWithinTimeSlot(filter.end, v.StartedAt, v.EndedAt)) {
			ok = false
		}

		if ok {
			ret = append(ret, v)
		}
	}

	return ret, nil
}

func (r *repoReservationImpl) CreateReservation(ctx context.Context, rv reservation) (reservation, error) {
	dbReservations = append(dbReservations, rv)
	return rv, nil
}

func (r *repoReservationImpl) CancelReservation(ctx context.Context, ID int) (reservation, error) {
	for _, v := range dbReservations {
		if v.ID == ID {
			v.IsCancelled = true
			return v, nil
		}
	}

	return reservation{}, nil
}

func (r *repoReservationImpl) GetReservation(ctx context.Context, ID int) (reservation, error) {
	for _, v := range dbReservations {
		if v.ID == ID {
			return v, nil
		}
	}

	return reservation{}, nil
}
