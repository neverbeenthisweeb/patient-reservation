package app

import "context"

var (
	dbPatients = []patient{
		{
			ID:        1,
			Name:      "Mr. Smith",
			CreatedAt: baseTime,
			UpdatedAt: baseTime,
		},
		{
			ID:        2,
			Name:      "Mrs. Ozora",
			CreatedAt: baseTime,
			UpdatedAt: baseTime,
		},
	}
)

type repoPatientImpl struct{}

func NewRepoPatientImpl() *repoPatientImpl {
	return &repoPatientImpl{}
}

func (r *repoPatientImpl) GetPatient(ctx context.Context, ID int) (patient, error) {
	for _, v := range dbPatients {
		if v.ID == ID {
			return v, nil
		}
	}

	return patient{}, nil
}
