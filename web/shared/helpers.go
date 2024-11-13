package shared

import (
	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	_sessions "github.com/zafir-co-ao/onna-narciso/internal/sessions"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/pages"
)

func CombineAppointmentsAndSessions(
	appointments []scheduling.AppointmentOutput,
	sessions []_sessions.SessionOutput,
) []pages.DailyAppointmentOptions {
	sessionMap := make(map[string]_sessions.SessionOutput)
	for _, session := range sessions {
		sessionMap[session.AppointmentID] = session
	}

	return xslices.Map(appointments, func(appointment scheduling.AppointmentOutput) pages.DailyAppointmentOptions {
		opts := pages.DailyAppointmentOptions{
			AppointmentID:     appointment.ID,
			AppointmentStatus: appointment.Status,
			AppointmentDate:   appointment.Date,
			AppointmentHour:   appointment.Hour,
			CustomerName:      appointment.CustomerName,
			ServiceName:       appointment.ServiceName,
			ProfessionalName:  appointment.ProfessionalName,
			SessionID:         "",
			SessionStatus:     "",
		}

		if session, found := sessionMap[appointment.ID]; found {
			opts.SessionID = session.ID
			opts.SessionStatus = session.Status
			return opts
		}

		return opts
	})
}
