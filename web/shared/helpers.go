package shared

import (
	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/session"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/pages"
)

func CombineAppointmentsAndSessions(appointments []scheduling.AppointmentOutput, sessions []session.SessionOutput) []pages.DailyAppointmentOpts {
	sessionMap := make(map[string]session.SessionOutput)
	for _, session := range sessions {
		sessionMap[session.AppointmentID] = session
	}

	return xslices.Map(appointments, func(a scheduling.AppointmentOutput) pages.DailyAppointmentOpts {
		if session, found := sessionMap[a.ID]; found {
			return pages.DailyAppointmentOpts{
				Appointment: a,
				Session:     session,
			}
		}

		return pages.DailyAppointmentOpts{Appointment: a, Session: session.SessionOutput{ID: ""}}
	})
}
