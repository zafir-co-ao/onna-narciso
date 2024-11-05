package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/components"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func NewAppointmentFinderHandler(f scheduling.AppointmentGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			_http.SendMethodNotAllowed(w)
			return
		}

		id := r.PathValue("id")

		o, err := f.Get(id)

		if errors.Is(scheduling.ErrAppointmentNotFound, err) {
			_http.SendNotFound(w, "Marcação não encontrada")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		_http.SendOk(w)
		components.AppointmentReschedulerDialog(o).Render(r.Context(), w)
	}
}
