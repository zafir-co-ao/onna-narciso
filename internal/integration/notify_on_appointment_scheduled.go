package integration

import (
	"fmt"
	"log/slog"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/notifications"
)

func NewNotifyOnAppointmentScheduledListener(s SchedulingServiceACL, crm CRMServiceACL, n notifications.Notifier) event.Handler {

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

		err = n.Notify(notifications.Contact{Mobile: c.PhoneNumber},
			notifications.Message{
				Subject: "Agendamento Realizado",
				Body:    fmt.Sprintf("Olá %s, seu agendamento foi realizado com sucesso! Obrigado pela preferência", c.Name),
			})

		if err != nil {
			slog.Error("Erro ao notificar cliente %s: %v", c.ID, err)
		}

	}

	return event.HandlerFunc(h)
}
