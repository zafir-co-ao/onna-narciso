package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/components"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleScheduleAppointment(s scheduling.AppointmentScheduler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			_http.SendMethodNotAllowed(w)
			return
		}

		duration, err := strconv.Atoi(r.FormValue("duration"))
		if err != nil {
			_http.SendBadRequest(w, "A duração da marcação está no formato inválido")
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

		o, err := s.Schedule(input)

		if errors.Is(scheduling.ErrCustomerNotFound, err) {
			_http.SendNotFound(w, "Cliente não encontrado")
			return
		}

		if errors.Is(scheduling.ErrCustomerRegistration, err) {
			_http.SendReponse(w, "Não foi possível registar o cliente", http.StatusInternalServerError)
			return
		}

		if errors.Is(scheduling.ErrProfessionalNotFound, err) {
			_http.SendNotFound(w, "Profissional não encontrado")
			return
		}

		if errors.Is(scheduling.ErrServiceNotFound, err) {
			_http.SendNotFound(w, "Serviço não encontrado")
			return
		}

		if errors.Is(scheduling.ErrBusyTime, err) {
			_http.SendBadRequest(w, "Horarário Indisponível")
			return
		}

		if errors.Is(scheduling.ErrInvalidDate, err) {
			_http.SendBadRequest(w, "A data para a marcação está no formato inválido")
			return
		}

		if errors.Is(scheduling.ErrInvalidHour, err) {
			_http.SendBadRequest(w, "A hora da marcação está no formato inválido")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		_http.SendCreated(w)
		components.Appointment(o, 6).Render(r.Context(), w)
	}
}
