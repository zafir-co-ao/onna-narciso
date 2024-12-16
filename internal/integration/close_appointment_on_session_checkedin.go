package integration

import (
	"log/slog"

	"github.com/kindalus/godx/pkg/event"
)

func NewCloseAppointmentOnSessionCheckedInListener(s SchedulingServiceACL) event.Handler {

	h := func(e event.Event) {
		p := e.Payload().(struct{ AppointmentID string })

		err := s.CloseAppointment(p.AppointmentID)
		if err != nil {
			slog.Error("Erro ao fechar o agendamento %s: %v", p.AppointmentID, err)
		}

	}

	return event.HandlerFunc(h)
}
