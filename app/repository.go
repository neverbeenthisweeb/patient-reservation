package app

import (
	"context"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	// baseTime simulates a default "now" value.
	baseTime time.Time = time.Now().Round(time.Hour)
)

type repo struct {
	slot        repoSlot
	patient     repoPatient
	doctor      repoDoctor
	reservation repoReservation
}

func NewRepo() *repo {
	return &repo{}
}

func (r *repo) SetSlot(rs repoSlot) {
	r.slot = rs
}

func (r *repo) SetPatient(rp repoPatient) {
	r.patient = rp
}

func (r *repo) SetDoctor(rd repoDoctor) {
	r.doctor = rd
}

func (r *repo) SetReservation(rr repoReservation) {
	r.reservation = rr
}

type repoSlot interface {
	GetSlots(ctx context.Context) ([]slot, error)
	GetSlot(ctx context.Context, ID int) (slot, error)
}

type repoPatient interface {
	GetPatient(ctx context.Context, ID int) (patient, error)
}

type repoDoctor interface {
	GetDoctor(ctx context.Context, ID int) (doctor, error)
}

type repoReservation interface {
	GetReservations(ctx context.Context, filter getReservationsFilter) ([]reservation, error)
	CreateReservation(ctx context.Context, rv reservation) (reservation, error)
	CancelReservation(ctx context.Context, ID int) (reservation, error)
	GetReservation(ctx context.Context, ID int) (reservation, error)
}

type getReservationsFilter struct {
	doctorID      int
	start         string
	end           string
	showCancelled bool
}
