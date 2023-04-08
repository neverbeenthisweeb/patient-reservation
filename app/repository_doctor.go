package app

import "context"

var (
	dbDoctors = []doctor{
		{
			ID:        1,
			Name:      "dr. Strange",
			CreatedAt: baseTime,
			UpdatedAt: baseTime,
		},
		{
			ID:        2,
			Name:      "dr. Boyke",
			CreatedAt: baseTime,
			UpdatedAt: baseTime,
		},
	}
)

type repoDoctorImpl struct{}

func NewRepoDoctorImpl() *repoDoctorImpl {
	return &repoDoctorImpl{}
}

func (r *repoDoctorImpl) GetDoctor(ctx context.Context, ID int) (doctor, error) {
	for _, v := range dbDoctors {
		if v.ID == ID {
			return v, nil
		}
	}

	return doctor{}, nil
}
