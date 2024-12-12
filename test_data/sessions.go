package testdata

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions"
)

var Sessions []sessions.Session = []sessions.Session{
	{
		ID:            nanoid.ID("1"),
		AppointmentID: nanoid.ID("4"),
		Status:        sessions.StatusStarted,
	},
	{
		ID:            nanoid.ID("2"),
		AppointmentID: nanoid.ID("7"),
		Status:        sessions.StatusStarted,
	},
}
