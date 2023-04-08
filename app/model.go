package app

import "time"

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
	IsCancelled bool      `json:"is_cancelled,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type slot struct {
	ID int `json:"id,omitempty"`
	// FIXME: Update StartedAt and EndedAt as string (e.g. 11:00 and 16.30)
	// FIXME: Need validation on string format?
	StartedAt string    `json:"started_at,omitempty"`
	EndedAt   string    `json:"ended_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// FIXME: Do unit test
// newSlotTime
//
// example:
//   in := "10:30"
//   out := newSlotTime(in)
//   fmt.Println(out) // FIXME: Update after writing test
func newSlotTime(s string) time.Time {
	panic("implement me")
}
