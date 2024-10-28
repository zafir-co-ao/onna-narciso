package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
)

func NewAppointmentCancelerHandler(u scheduling.AppointmentCanceler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			sendMethodNotAllowed(w)
			return
		}

		id := r.PathValue("id")

		err := u.Execute(id)

		if errors.Is(scheduling.ErrInvalidStatusToCancel, err) {
			sendBadRequest(w, "Estado inválido para cancelar")
			return
		}

		if errors.Is(scheduling.ErrAppointmentNotFound, err) {
			sendBadRequest(w, "Marcação não encontrada")
			return
		}

		if !errors.Is(nil, err) {
			sendServerError(w)
			return
		}

		sendOk(w)
	}
}
