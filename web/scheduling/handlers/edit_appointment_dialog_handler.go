package handlers

import (
	"errors"
	"net/http"
	"slices"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
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

		services := slices.Clone(testdata.Services)

		professionals := make([]scheduling.Professional, 0, 0)
		// TODO - Utilizar o repositório de profissionais para encontrar os serviços com base no profissional
		for _, p := range testdata.Professionals {
			if slices.Contains(p.ServicesIDS, nanoid.ID(o.ServiceID)) {
				professionals = append(professionals, p)
			}
		}

		opts := components.AppointmentReschedulerOptions{
			Appointment:   o,
			Professionals: professionals,
			Services:      services,
		}

		_http.SendOk(w)
		components.AppointmentReschedulerDialog(opts).Render(r.Context(), w)
	}
}
