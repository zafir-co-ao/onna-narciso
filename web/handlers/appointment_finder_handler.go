package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/web/components"
)

func NewAppointmentFinderHandler(f scheduling.AppointmentFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			sendMethodNotAllowed(w)
			return
		}

		id := r.PathValue("id")

		o, err := f.Execute(id)

		if errors.Is(scheduling.ErrAppointmentNotFound, err) {
			sendNotFound(w, "Marcação não encontrada")
			return
		}

		if !errors.Is(nil, err) {
			sendServerError(w)
			return
		}

		sendOk(w)
		components.RescheduleEventForm(o).Render(r.Context(), w)

	}
}
