package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/web/components"
)

func NewAppointmentFinderHandler(f scheduling.AppointmentGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			sendMethodNotAllowed(w)
			return
		}

		id := r.PathValue("id")

		o, err := f.Get(id)

		if errors.Is(scheduling.ErrAppointmentNotFound, err) {
			sendNotFound(w, "Marcação não encontrada")
			return
		}

		if !errors.Is(nil, err) {
			sendServerError(w)
			return
		}

		sendOk(w)
		components.AppointmentReschedulingForm(o).Render(r.Context(), w)

	}
}
