package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/web/components"
)

func NewAppointmentSchedulerHandler(u scheduling.AppointmentScheduler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			sendMethodNotAllowed(w)
			return
		}

		duration, err := strconv.Atoi(r.FormValue("duration"))
		if err != nil {
			sendBadRequest(w, "A duração da marcação está no formato inválido")
			return
		}

		input := scheduling.AppointmentSchedulerInput{
			ProfessionalID: r.Form.Get("professional-id"),
			ServiceID:      r.Form.Get("service-id"),
			Date:           r.Form.Get("date"),
			StartHour:      r.Form.Get("start"),
			CustomerID:     r.Form.Get("customer-id"),
			CustomerName:   r.Form.Get("customer-name"),
			CustomerPhone:  r.Form.Get("customer-phone"),
			Duration:       duration,
		}

		o, err := u.Schedule(input)

		if errors.Is(scheduling.ErrCustomerNotFound, err) {
			sendNotFound(w, "Cliente não encontrado")
			return
		}

		if errors.Is(scheduling.ErrCustomerRegistration, err) {
			sendReponse(w, "Não foi possível registar o cliente", http.StatusInternalServerError)
			return
		}

		if errors.Is(scheduling.ErrProfessionalNotFound, err) {
			sendNotFound(w, "Profissional não encontrado")
			return
		}

		if errors.Is(scheduling.ErrServiceNotFound, err) {
			sendNotFound(w, "Serviço não encontrado")
			return
		}

		if errors.Is(scheduling.ErrBusyTime, err) {
			sendBadRequest(w, "Horarário Indisponível")
			return
		}

		if errors.Is(scheduling.ErrInvalidDate, err) {
			sendBadRequest(w, "A data para a marcação está no formato inválido")
			return
		}

		if errors.Is(scheduling.ErrInvalidHour, err) {
			sendBadRequest(w, "A hora da marcação está no formato inválido")
			return
		}

		if !errors.Is(nil, err) {
			sendServerError(w)
			return
		}

		sendCreated(w)
		components.Event(o, 8).Render(r.Context(), w)
	}
}
