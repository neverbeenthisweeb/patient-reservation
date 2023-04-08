package app

import (
	"context"
)

var (
	dbSlots []slot = []slot{
		{
			ID:        1,
			StartedAt: "10:00",
			EndedAt:   "12:00",
			CreatedAt: baseTime,
			UpdatedAt: baseTime,
		},
		{
			ID:        2,
			StartedAt: "12:00",
			EndedAt:   "14:00",
			CreatedAt: baseTime,
			UpdatedAt: baseTime,
		},
		{
			ID:        3,
			StartedAt: "14:00",
			EndedAt:   "16:00",
			CreatedAt: baseTime,
			UpdatedAt: baseTime,
		},
	}
)

type repoSlotImpl struct{}

func NewRepoSlotImpl() *repoSlotImpl {
	return &repoSlotImpl{}
}

func (r *repoSlotImpl) GetSlots(ctx context.Context) ([]slot, error) {
	return dbSlots, nil
}

func (r *repoSlotImpl) GetSlot(ctx context.Context, ID int) (slot, error) {
	for _, v := range dbSlots {
		if v.ID == ID {
			return v, nil
		}
	}

	return slot{}, nil
}
