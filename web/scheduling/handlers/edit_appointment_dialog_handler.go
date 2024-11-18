package handlers

import (
	"errors"
	"net/http"
	"slices"

	"github.com/kindalus/godx/pkg/coalesce"
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

		serviceID := coalesce.Fallback(r.FormValue("service-id"), o.ServiceID)
		professionalID := coalesce.Fallback(r.FormValue("professional-id"), o.ProfessionalID)

		services := slices.Clone(testdata.Services)

		professionals := make([]scheduling.Professional, 0, 0)
		// TODO - Utilizar o repositório de profissionais para encontrar os serviços com base no profissional
		for _, p := range testdata.Professionals {
			if slices.Contains(p.ServicesIDS, nanoid.ID(serviceID)) {
				professionals = append(professionals, p)
			}
		}

		constainsCurrentProfessional := func(p scheduling.Professional) bool {
			return p.ID.String() == professionalID
		}

		if !slices.ContainsFunc(professionals, constainsCurrentProfessional) {
			professionalID = professionals[0].ID.String()
		}

		// TODO - Usar o repositório dos serviços para filtrar os serviços

		opts := components.AppointmentReschedulerOptions{
			Appointment:          o,
			Professionals:        professionals,
			SelectedProfessional: professionalID,
			SelectedService:      serviceID,
			Services:             services,
			HandlerURL:           r.URL.Path,
		}

		_http.SendOk(w)
		components.AppointmentReschedulerDialog(opts).Render(r.Context(), w)
	}
}
