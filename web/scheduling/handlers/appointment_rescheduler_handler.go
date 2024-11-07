package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"

	"github.com/zafir-co-ao/onna-narciso/web/scheduling/components"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleRescheduleNewAppointment(re scheduling.AppointmentRescheduler) func(w http.ResponseWriter, r *http.Request) {
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

		input := scheduling.AppointmentReschedulerInput{
			ID:        r.Form.Get("id"),
			Date:      r.Form.Get("date"),
			StartHour: r.Form.Get("start"),
			Duration:  duration,
		}

		o, err := re.Reschedule(input)
		if errors.Is(scheduling.ErrInvalidStatusToReschedule, err) {
			_http.SendBadRequest(w, "Estado inválido para reagendar")
			return
		}

		if errors.Is(scheduling.ErrBusyTime, err) {
			_http.SendBadRequest(w, "Horário Indisponível")
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

		if errors.Is(scheduling.ErrAppointmentNotFound, err) {
			_http.SendNotFound(w, "Marcação não encontrada")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		_http.SendOk(w)
		components.Appointment(o, 6).Render(r.Context(), w)
	}
}
