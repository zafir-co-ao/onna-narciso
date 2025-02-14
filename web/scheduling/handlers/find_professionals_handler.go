package handlers

import (
	"errors"
	"net/http"
	"slices"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/hr"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
	"github.com/zafir-co-ao/onna-narciso/web/shared/components"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleFindProfessionals(pf hr.ProfessionalFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		serviceID := r.FormValue("service-id")

		professionalOutput, err := pf.FindAll()

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		servicesIDs := make([]nanoid.ID, 0)
		for _, p := range professionalOutput {
			for _, s := range p.Services {
				servicesIDs = append(servicesIDs, nanoid.ID(s.ID))
			}
		}

		professionals := xslices.Map(professionalOutput, func(p hr.ProfessionalOutput) scheduling.Professional {
			return scheduling.Professional{
				ID:          nanoid.ID(p.ID),
				Name:        name.Name(p.Name),
				ServicesIDs: servicesIDs,
			}
		})

		professionalsOptions := make([]scheduling.Professional, 0)
		for _, p := range professionals {
			if slices.Contains(p.ServicesIDs, nanoid.ID(serviceID)) {
				professionalsOptions = append(professionalsOptions, p)
			}
		}

		components.Dropdown(
			"professional-id",
			professionals[0].ID.String(),
			components.WithOptions(components.ProfessionalsToOptions(professionalsOptions)...),
		).Render(r.Context(), w)
	}
}
