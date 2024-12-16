package integration

import (
	"fmt"
	"log/slog"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/notifications"
)

func NewNotifyOnAppointmentRescheduledListener(s SchedulingServiceACL, crm CRMServiceACL, n notifications.Notifier) event.Handler {

	h := func(e event.Event) {
		id := e.Header(event.HeaderAggregateID)

		a, err := s.FindAppointmentByID(nanoid.ID(id))
		if err != nil {
			slog.Error("Erro ao carregar agendamento %s: %v", id, err)
			return
		}

		c, err := crm.GetCustomer(a.CustomerID)
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

	return event.HandlerFunc(h)
}
