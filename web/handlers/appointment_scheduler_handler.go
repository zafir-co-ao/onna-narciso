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
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		duration, err := strconv.ParseInt(r.FormValue("duration"), 10, 64)
		if err != nil {
			http.Error(w, "A duração da marcação está no formato inválido", http.StatusBadRequest)
			return
		}

		input := scheduling.AppointmentSchedulerInput{
			ProfessionalID: r.Form.Get("professional_id"),
			ServiceID:      r.Form.Get("service_id"),
			Date:           r.Form.Get("date"),
			StartHour:      r.Form.Get("start"),
			CustomerID:     r.Form.Get("customer_id"),
			CustomerName:   r.Form.Get("customer_name"),
			CustomerPhone:  r.Form.Get("customer_phone"),
			Duration:       int(duration),
		}

		o, err := u.Schedule(input)

		if errors.Is(scheduling.ErrCustomerNotFound, err) {
			http.Error(w, "Cliente não encontrado", http.StatusNotFound)
			return
		}

		if errors.Is(scheduling.ErrCustomerRegistration, err) {
			http.Error(w, "Não foi possível registar o cliente", http.StatusInternalServerError)
			return
		}

		if errors.Is(scheduling.ErrProfessionalNotFound, err) {
			http.Error(w, "Profissional não encontrado", http.StatusNotFound)
			return
		}

		if errors.Is(scheduling.ErrServiceNotFound, err) {
			http.Error(w, "Serviço não encontrado", http.StatusNotFound)
			return
		}

		if errors.Is(scheduling.ErrBusyTime, err) {
			http.Error(w, "Horarário Indisponível", http.StatusBadRequest)
			return
		}

		if errors.Is(scheduling.ErrInvalidDate, err) {
			http.Error(w, "A data para a marcação está no formato inválido", http.StatusBadRequest)
			return
		}

		if errors.Is(scheduling.ErrInvalidHour, err) {
			http.Error(w, "A hora da marcação está no formato inválido", http.StatusBadRequest)
			return
		}

		if !errors.Is(nil, err) {
			http.Error(w, "Erro desconhecido, contacte o Administrador", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		components.Event(o, 8).Render(r.Context(), w)
	}
}
