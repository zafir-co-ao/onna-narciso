package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/web/hr/components"
)

func HandleCreateProfessionalDialog(u services.ServiceFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.FormValue("hx-post")

		// o, err := u.FindAll()

		// if !errors.Is(nil, err) {
		// 	_http.SendServerError(w)
		// 	return
		// }

		o := []services.ServiceOutput{
			{ID: "1", Name: "Manicure"},
			{ID: "2", Name: "Pedicure"},
		}

		components.ProfessionalCreateDialog(url, o).Render(r.Context(), w)
	}
}
