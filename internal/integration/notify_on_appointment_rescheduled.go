package integration

import (
	"fmt"
	"log/slog"

	"github.com/kindalus/godx/pkg/event"
	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/notifications"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
)

func ListenAndNotifyOnAppointmentRescheduled(bus event.Bus, n notifications.Notifier, afinder scheduling.AppointmentFinder, cfinder crm.CustomerFinder) {

	h := func(e event.Event) {
		id := e.Header(event.HeaderAggregateID)

		a, err := afinder.FindByID(id)
		if err != nil {
			slog.Error("Erro ao carregar agendamento %s: %v", id, err)
			return
		}

		c, err := cfinder.FindByID(a.CustomerID)
		if err != nil {
			slog.Error("Erro ao carregar cliente %s: %v", a.CustomerID, err)
			return
		}

		err = n.Notify(
			notifications.Contact{Mobile: c.PhoneNumber},
			notifications.Message{
				Subject: "Agendamento Reagendado",
				Body:    fmt.Sprintf("Olá %s, seu agendamento foi reagendado com sucesso! Obrigado pela preferência", c.Name),
			})

		if err != nil {
			slog.Error("Erro ao notificar cliente %s: %v", a.CustomerID, err)
		}

	}

	bus.Subscribe(scheduling.EventAppointmentScheduled, event.HandlerFunc(h))
}
