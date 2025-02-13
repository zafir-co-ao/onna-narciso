package handlers

import (
	"errors"
	"net/http"
	"slices"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/hr"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/duration"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/components"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleEditAppointmentDialog(af scheduling.AppointmentFinder, sf services.ServiceFinder, pf hr.ProfessionalFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		o, err := af.FindByID(r.PathValue("id"))

		if errors.Is(err, scheduling.ErrAppointmentNotFound) {
			_http.SendNotFound(w, "Marcação não encontrada")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		professionalOutput, err := pf.FindAll()

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		servicesIDs := make([]string, 0)
		for _, p := range professionalOutput {
			for _, s := range p.Services {
				servicesIDs = append(servicesIDs, s.ID)
			}
		}

		professionals := xslices.Map(professionalOutput, func(p hr.ProfessionalOutput) scheduling.Professional {
			ids := xslices.Map(servicesIDs, func(id string) nanoid.ID {
				return nanoid.ID(id)
			})

			return scheduling.Professional{
				ID:          nanoid.ID(p.ID),
				Name:        name.Name(p.Name),
				ServicesIDs: ids,
			}
		})

		professionalsOptions := make([]scheduling.Professional, 0)
		for _, p := range professionals {
			if slices.Contains(p.ServicesIDs, nanoid.ID(o.ServiceID)) {
				professionalsOptions = append(professionalsOptions, p)
			}
		}

		ids := getServicesIDs(professionalsOptions)
		servicesOutput, err := sf.FindByIDs(ids)

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		servicesOptions := xslices.Map(servicesOutput, func(s services.ServiceOutput) scheduling.Service {
			return scheduling.Service{
				ID:       nanoid.ID(s.ID),
				Name:     name.Name(s.Name),
				Duration: duration.Duration(s.Duration),
			}
		})

		opts := components.AppointmentReschedulerOptions{
			Appointment:   o,
			Professionals: professionalsOptions,
			Services:      servicesOptions,
		}

		_http.SendOk(w)
		components.AppointmentReschedulerDialog(opts).Render(r.Context(), w)
	}
}

func getServicesIDs(p []scheduling.Professional) []string {
	sid := make(map[string]struct{})
	var uniqueIDs []string

	for _, p := range p {
		for _, id := range p.ServicesIDs {
			sid[id.String()] = struct{}{}

			if _, exists := sid[id.String()]; !exists {
				uniqueIDs = append(uniqueIDs, id.String())
			}
		}
	}

	for id := range sid {
		uniqueIDs = append(uniqueIDs, id)
	}

	return uniqueIDs
}
