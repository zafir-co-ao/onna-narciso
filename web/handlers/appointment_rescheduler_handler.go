package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/web/components"
)

func NewAppointmentReschedulerHandler(re scheduling.AppointmentRescheduler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if http.MethodPost != r.Method {
			sendMethodNotAllowed(w)
			return
		}

		duration, err := strconv.Atoi(r.FormValue("duration"))
		if err != nil {
			sendBadRequest(w, "A duração da marcação está no formato inválido")
			return
		}

		i := scheduling.AppointmentReschedulerInput{
			ID:        r.Form.Get("id"),
			Date:      r.Form.Get("date"),
			StartHour: r.Form.Get("start"),
			Duration:  duration,
		}

		o, err := re.Execute(i)
		if errors.Is(scheduling.ErrInvalidStatusToReschedule, err) {
			sendBadRequest(w, "Estado inválido para reagendar")
			return
		}

		if errors.Is(scheduling.ErrBusyTime, err) {
			sendBadRequest(w, "Horário Indisponível")
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

		if errors.Is(scheduling.ErrAppointmentNotFound, err) {
			sendNotFound(w, "Marcação não encontrada")
			return
		}

		if !errors.Is(nil, err) {
			sendServerError(w)
			return
		}

		sendOk(w)
		components.Event(o, 8).Render(r.Context(), w)
	}
}
