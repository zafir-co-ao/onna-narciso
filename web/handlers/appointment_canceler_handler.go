package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
)

func NewAppointmentCancelerHandler(u scheduling.AppointmentCanceler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		id := r.PathValue("id")

		err := u.Execute(id)

		if errors.Is(scheduling.ErrAppointmentNotFound, err) {
			http.Error(w, "appointment not found", http.StatusNotFound)
			return
		}

		if errors.Is(scheduling.ErrInvalidStatusToCancel, err) {
			http.Error(w, "invalid status to cancel", http.StatusBadRequest)
			return
		}

		if !errors.Is(nil, err) {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("appointment cancelled"))
	}
}
