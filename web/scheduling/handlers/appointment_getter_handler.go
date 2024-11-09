package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/components"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleEditAppointmentDialog(g scheduling.AppointmentGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		o, err := g.Get(r.PathValue("id"))

		if errors.Is(err, scheduling.ErrAppointmentNotFound) {
			_http.SendNotFound(w, "Marcação não encontrada")
			return
		}

		if err != nil {
			_http.SendServerError(w)
			return
		}

		target := r.FormValue("hx-target")
		swap := r.FormValue("hx-swap")

		_http.SendOk(w)
		components.AppointmentReschedulerDialog(o, target, swap).Render(r.Context(), w)
	}
}
