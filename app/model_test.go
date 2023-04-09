package app

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewSlotTime(t *testing.T) {
	now := time.Now().UTC()

	cases := []struct {
		in  string
		out time.Time
	}{
		{
			in:  "10:30",
			out: time.Date(now.Year(), now.Month(), now.Day(), 10, 30, 0, 0, now.Location()),
		},
		{
			in:  "16:15",
			out: time.Date(now.Year(), now.Month(), now.Day(), 16, 15, 0, 0, now.Location()),
		},
	}

	for _, tc := range cases {
		st := newSlotTime(tc.in)
		require.Equal(t, st, tc.out)
	}
}

func TestIsWithinTimeSlot(t *testing.T) {
	now := time.Now().UTC()

	cases := []struct {
		s     string
		start time.Time
		end   time.Time
		ok    bool
	}{
		{
			s:     "10:30",
			start: time.Date(now.Year(), now.Month(), now.Day(), 10, 45, 0, 0, now.Location()),
			end:   time.Date(now.Year(), now.Month(), now.Day(), 11, 00, 0, 0, now.Location()),
			ok:    false,
		},
		{
			s:     "10:30",
			start: time.Date(now.Year(), now.Month(), now.Day(), 10, 15, 0, 0, now.Location()),
			end:   time.Date(now.Year(), now.Month(), now.Day(), 10, 45, 0, 0, now.Location()),
			ok:    true,
		},
	}

	for _, tc := range cases {
		ok := isWithinTimeSlot(tc.s, tc.start, tc.end)
		require.Equal(t, ok, tc.ok)
	}
}
