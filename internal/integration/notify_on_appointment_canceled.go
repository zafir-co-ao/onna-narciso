package integration

import (
	"fmt"
	"log/slog"

	"github.com/kindalus/godx/pkg/event"
	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/notifications"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
)

func ListenAndNotifyOnAppointmentCanceled(bus event.Bus, n notifications.Notifier, finder crm.CustomerFinder) {
	h := func(e event.Event) {
		id := e.Header(event.HeaderAggregateID)

		c, err := finder.FindByID(id)
		if err != nil {
			slog.Error("Erro ao carregar cliente %s: %v", id, err)
			return
		}

		err = n.Notify(notifications.Contact{Mobile: c.PhoneNumber},
			notifications.Message{
				Subject: "Agendamento cancelado",
				Body:    fmt.Sprintf("Olá %s, seu agendamento foi cancelado.", c.Name),
			})

		if err != nil {
			slog.Error("Erro ao notificar cliente %s: %v", c.ID, err)
		}
	}

	bus.Subscribe(scheduling.EventAppointmentCanceled, event.HandlerFunc(h))
}
