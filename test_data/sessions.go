package testdata

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/session"
)

var Sessions []session.Session = []session.Session{
	{
		ID:            nanoid.ID("1"),
		AppointmentID: nanoid.ID("4"),
		Status:        session.StatusStarted,
	},
}
