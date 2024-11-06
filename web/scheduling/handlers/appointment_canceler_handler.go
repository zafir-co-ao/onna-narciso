package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func NewAppointmentCancelerHandler(c scheduling.AppointmentCanceler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			_http.SendMethodNotAllowed(w)
			return
		}

		err := c.Execute(r.PathValue("id"))

		if errors.Is(scheduling.ErrInvalidStatusToCancel, err) {
			_http.SendBadRequest(w, "Estado inválido para cancelar")
			return
		}

		if errors.Is(scheduling.ErrAppointmentNotFound, err) {
			_http.SendBadRequest(w, "Marcação não encontrada")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		_http.SendOk(w)
	}
}
