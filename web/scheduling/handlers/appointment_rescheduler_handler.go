package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/hour"

	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleRescheduleAppointment(re scheduling.AppointmentRescheduler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		duration, err := strconv.Atoi(r.FormValue("duration"))
		if err != nil {
			_http.SendBadRequest(w, "A duração da marcação está no formato inválido")
			return
		}

		i := scheduling.AppointmentReschedulerInput{
			ID:             r.FormValue("id"),
			Date:           r.FormValue("date"),
			Hour:           r.FormValue("hour"),
			ProfessionalID: r.FormValue("professional-id"),
			ServiceID:      r.FormValue("service-id"),
			Duration:       duration,
		}

		err = re.Reschedule(i)
		if errors.Is(err, scheduling.ErrInvalidStatusToReschedule) {
			_http.SendBadRequest(w, "Estado inválido para reagendar")
			return
		}

		if errors.Is(err, scheduling.ErrBusyTime) {
			_http.SendBadRequest(w, "Horário Indisponível")
			return
		}

		if errors.Is(err, date.ErrInvalidFormat) {
			_http.SendBadRequest(w, "A data para a marcação está no formato inválido")
			return
		}

		if errors.Is(err, date.ErrDateInPast) {
			_http.SendBadRequest(w, "A marcação não pode ser feita para uma data no passado")
			return
		}

		if errors.Is(err, hour.ErrInvalidFormat) {
			_http.SendBadRequest(w, "A hora da marcação está no formato inválido")
			return
		}

		if errors.Is(err, scheduling.ErrInvalidService) {
			_http.SendBadRequest(w, "Indisponibilidade do serviço para o profissional")
			return
		}

		if errors.Is(err, scheduling.ErrAppointmentNotFound) {
			_http.SendNotFound(w, "Marcação não encontrada")
			return
		}

		if errors.Is(err, scheduling.ErrProfessionalNotFound) {
			_http.SendNotFound(w, "Profissional não encontrado")
			return
		}

		if errors.Is(err, scheduling.ErrServiceNotFound) {
			_http.SendNotFound(w, "Serviço não encontrado")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		w.Header().Set("X-Reload-Page", "ReloadPage")
		_http.SendOk(w)
	}
}
