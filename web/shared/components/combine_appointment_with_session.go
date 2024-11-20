package components

import (
	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	_sessions "github.com/zafir-co-ao/onna-narciso/internal/sessions"
)

type DailyAppointmentOptions struct {
	AppointmentID     string
	AppointmentStatus string
	AppointmentDate   string
	AppointmentHour   string
	CustomerName      string
	ServiceName       string
	ProfessionalName  string
	SessionID         string
	SessionStatus     string
}

func CombineAppointmentsWithSessions(a []scheduling.AppointmentOutput, s []_sessions.SessionOutput) []DailyAppointmentOptions {
	sessionMap := make(map[string]_sessions.SessionOutput)
	for _, session := range s {
		sessionMap[session.AppointmentID] = session
	}

	return xslices.Map(a, func(a scheduling.AppointmentOutput) DailyAppointmentOptions {
		opts := DailyAppointmentOptions{
			AppointmentID:     a.ID,
			AppointmentStatus: a.Status,
			AppointmentDate:   a.Date,
			AppointmentHour:   a.Hour,
			CustomerName:      a.CustomerName,
			ServiceName:       a.ServiceName,
			ProfessionalName:  a.ProfessionalName,
			SessionID:         "",
			SessionStatus:     "",
		}

		if session, found := sessionMap[a.ID]; found {
			opts.SessionID = session.ID
			opts.SessionStatus = session.Status
			return opts
		}

		return opts
	})
}
