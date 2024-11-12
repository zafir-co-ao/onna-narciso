package shared

import (
	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/session"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/pages"
)

func CombineAppointmentsAndSessions(appointments []scheduling.AppointmentOutput, sessions []session.SessionOutput) []pages.DailyAppointmentOptions {
	sessionMap := make(map[string]session.SessionOutput)
	for _, session := range sessions {
		sessionMap[session.AppointmentID] = session
	}

	return xslices.Map(appointments, func(appointment scheduling.AppointmentOutput) pages.DailyAppointmentOptions {
		if session, found := sessionMap[appointment.ID]; found {
			return pages.DailyAppointmentOptions{
				Appointment: appointment,
				Session:     session,
			}
		}

		return pages.DailyAppointmentOptions{Appointment: appointment, Session: session.SessionOutput{ID: ""}}
	})
}
