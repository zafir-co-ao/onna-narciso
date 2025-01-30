package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/web/hr/components"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleCreateProfessionalDialog(u services.ServiceFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.FormValue("hx-post")

		o, err := u.FindAll()

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		_http.SendOk(w)
		components.ProfessionalCreateDialog(url, o).Render(r.Context(), w)
	}
}
