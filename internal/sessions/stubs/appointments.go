package stubs

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions"
	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
)

func NewSchedulingServiceACL() sessions.SchedulingServiceACLFunc {
	var aacl sessions.SchedulingServiceACLFunc = func(id nanoid.ID) (sessions.Appointment, error) {
		for _, a := range testdata.Appointments {
			if a.ID == id {
				return sessions.Appointment{
					ID:               a.ID,
					ProfessionalID:   a.ProfessionalID,
					CustomerID:       a.CustomerID,
					ServiceID:        a.ServiceID,
					CustomerName:     string(a.CustomerName),
					ProfessionalName: string(a.ProfessionalName),
					ServiceName:      string(a.ServiceName),
					Closed:           a.IsClosed(),
					Canceled:         a.IsCancelled(),
					Date:             a.Date,
				}, nil
			}
		}

		return sessions.Appointment{}, sessions.ErrAppointmentNotFound
	}

	return aacl
}
