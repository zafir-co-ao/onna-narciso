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

		professionalID := r.FormValue("professional-id")
		if professionalID == "" {
			professionalID = o.ProfessionalID
		}

		var services []scheduling.Service
		var servicesIDs []nanoid.ID

		// TODO - Utilizar o repositório de profissionais para encontrar os serviços com base no profissional
		for _, p := range testdata.Professionals {
			if p.ID.String() == professionalID {
				servicesIDs = p.ServicesIDS
			}
		}

		// TODO - Usar o repositório dos serviços para filtrar os serviços
		for _, s := range testdata.Services {
			if slices.Contains(servicesIDs, s.ID) {
				services = append(services, s)
			}
		}

		opts := components.AppointmentReschedulerOptions{
			Appointment:          o,
			Professionals:        testdata.Professionals,
			SelectedProfessional: professionalID,
			SelectedService:      services[0].ID.String(),
			Services:             services,
			HxTarget:             r.FormValue("hx-target"),
			HxSwap:               r.FormValue("hx-swap"),
		}

		_http.SendOk(w)
		components.AppointmentReschedulerDialog(opts).Render(r.Context(), w)
	}
}
