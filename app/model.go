package app

import (
	"strconv"
	"strings"
	"time"
)

type patient struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type doctor struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type reservation struct {
	ID          int       `json:"id,omitempty"`
	PatientID   int       `json:"patient_id,omitempty"`
	DoctorID    int       `json:"doctor_id,omitempty"`
	StartedAt   time.Time `json:"started_at,omitempty"`
	EndedAt     time.Time `json:"ended_at,omitempty"`
	IsCancelled bool      `json:"is_cancelled"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type slot struct {
	ID        int       `json:"id,omitempty"`
	StartedAt string    `json:"started_at,omitempty"`
	EndedAt   string    `json:"ended_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// newSlotTime
//
// example:
//   in := "10:30"
//   out := newSlotTime(in)
//   fmt.Println(out) // 2023-04-09 10:30:00 +0000 UTC
func newSlotTime(s string) time.Time {
	ss := strings.Split(s, ":")
	hh, _ := strconv.Atoi(ss[0])
	mm, _ := strconv.Atoi(ss[1])
	now := time.Now().UTC()
	return time.Date(now.Year(), now.Month(), now.Day(), hh, mm, 0, 0, now.Location())
}

func isWithinTimeSlot(s string, start, end time.Time) bool {
	slotTime := newSlotTime(s)

	if slotTime.Equal(start) || slotTime.Equal(end) {
		return true
	}

	return slotTime.After(start) && slotTime.Before(end)
}
