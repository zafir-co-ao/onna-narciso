package handlers

import (
	"net/http"
	"slices"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
	"github.com/zafir-co-ao/onna-narciso/web/shared/components"
)

func HandleFindProfessionals() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		serviceID := r.FormValue("service-id")

		professionals := make([]scheduling.Professional, 0, 0)
		// TODO - Utilizar o repositório de profissionais para encontrar os serviços com base no profissional
		for _, p := range testdata.Professionals {
			if slices.Contains(p.ServicesIDS, nanoid.ID(serviceID)) {
				professionals = append(professionals, p)
			}
		}

		components.Dropdown(
			"professional-id",
			professionals[0].ID.String(),
			components.WithOptions(components.ProfessionalsToOptions(professionals)...),
		).Render(r.Context(), w)
	}
}
