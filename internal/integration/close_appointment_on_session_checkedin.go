package integration

import (
	"log/slog"

	"github.com/kindalus/godx/pkg/event"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions"
)

func ListenAndCloseAppointmentOnSessionCheckedIn(bus event.Bus, closer scheduling.AppointmentCloser) {

	h := func(e event.Event) {
		p := e.Payload().(struct{ AppointmentID string })

		err := closer.Close(p.AppointmentID)
		if err != nil {
			slog.Error("Erro ao fechar o agendamento %s: %v", p.AppointmentID, err)
		}

	}

	bus.Subscribe(sessions.EventSessionCheckedIn, event.HandlerFunc(h))
}
