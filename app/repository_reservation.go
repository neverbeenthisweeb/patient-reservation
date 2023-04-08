package app

import "context"

const (
	// Queue counter must obey this.
	initialReservationID = 1
)

var (
	dbReservations = []reservation{
		{
			ID:        1,
			PatientID: 1,
			DoctorID:  1,
			// StartedAt: ,
			// EndedAt: ,
			IsCancelled: false,
			CreatedAt:   baseTime,
			UpdatedAt:   baseTime,
		},
		{
			ID:        2,
			PatientID: 2,
			DoctorID:  1,
			// StartedAt: ,
			// EndedAt: ,
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
		if filter.doctorID != 0 && filter.doctorID == v.DoctorID {
			ret = append(ret, v)
		}
	}

	return ret, nil
}

func (r *repoReservationImpl) CreateReservation(ctx context.Context, rv reservation) (reservation, error) {
	dbReservations = append(dbReservations, rv)
	return rv, nil
}
