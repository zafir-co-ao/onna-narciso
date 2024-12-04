package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleCancelAppointment(c scheduling.AppointmentCanceler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := c.Cancel(r.PathValue("id"))

		if errors.Is(err, scheduling.ErrInvalidStatusToCancel) {
			_http.SendBadRequest(w, "Estado inválido para cancelar")
			return
		}

		if errors.Is(err, scheduling.ErrAppointmentNotFound) {
			_http.SendBadRequest(w, "Marcação não encontrada")
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
