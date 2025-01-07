package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/hr"
	"github.com/zafir-co-ao/onna-narciso/web/hr/pages"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleFindProfessionals(u hr.ProfessionalFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// o, err := u.FindAll()

		// if !errors.Is(nil, err) {
		// 	_http.SendServerError(w)
		// 	return
		// }

		o := []hr.ProfessionalOutput{
			{
				ID:   "1",
				Name: "Jonathan Paulino",
				Services: []hr.ServiceOutput{
					{ID: "1", Name: "Manicure"},
					{ID: "2", Name: "Pedicure"},
					{ID: "3", Name: "Massagem"},
				},
			},
			{
				ID:   "2",
				Name: "Kevin de Bruine",
				Services: []hr.ServiceOutput{
					{ID: "1", Name: "Manicure"},
					{ID: "2", Name: "Massagem"},
				},
			},
			{
				ID:   "3",
				Name: "Luana Targinho",
				Services: []hr.ServiceOutput{
					{ID: "1", Name: "Pedicure"},
					{ID: "2", Name: "Manicure"},
				},
			},
		}

		_http.SendOk(w)
		pages.ListProfessionals(o).Render(r.Context(), w)
	}
}
