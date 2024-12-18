package integration

import (
	"fmt"
	"log/slog"

	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/notifications"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
)

func SendAppointmentConfirmationRequest(n notifications.Notifier, arepo scheduling.AppointmentRepository, cfinder crm.CustomerFinder) {

	appointments, err := arepo.FindByDate(date.Today().AddDate(0, 0, -1))
	if err != nil {
		slog.Error("Erro ao carregar agendamentos para %s: %v", "hoje", err)
		return
	}

	for _, a := range appointments {

		c, err := cfinder.FindByID(a.CustomerID.String())
		if err != nil {
			slog.Error("Erro ao carregar cliente %s: %v", a.CustomerID.String(), err)
			return
		}

		go func() {
			err := n.Notify(notifications.Contact{Mobile: c.PhoneNumber},
				notifications.Message{
					Subject: "Confirmação de agendamento",
					Body:    fmt.Sprintf("Olá %s, confirme seu agendamento para amanhã.", c.Name),
				})

			if err != nil {
				slog.Error("Erro ao notificar cliente %s: %v", c.ID, err)
			}
		}()

	}
}
