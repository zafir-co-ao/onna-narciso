package integration

import (
	"fmt"
	"log/slog"

	"github.com/kindalus/godx/pkg/event"
	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/notifications"
)

func ListAndSendWelcomeToOnna(bus event.Bus, n notifications.Notifier, finder crm.CustomerFinder) {
	h := func(e event.Event) {
		id := e.Header(event.HeaderAggregateID)

		c, err := finder.FindByID(id)
		if err != nil {
			slog.Error("Erro ao carregar cliente %s: %v", id, err)
			return
		}

		err = n.Notify(notifications.Contact{Mobile: c.PhoneNumber},
			notifications.Message{
				Subject: "Bem-vindo à Onna",
				Body:    fmt.Sprintf("Olá %s, seja bem-vindo à Onna!", c.Name),
			})

		if err != nil {
			slog.Error("Erro ao notificar cliente %s: %v", c.ID, err)
		}
	}

	bus.Subscribe(crm.EventCustomerCreated, event.HandlerFunc(h))
}
