package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/web/services/components"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleEditServiceDialog(u services.ServiceGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.FormValue("id")
		url := r.FormValue("hx-put")

		s, err := u.Get(id)
		if !errors.Is(nil, err) {
			_http.SendNotFound(w, "Serviço não encontrado")
		}

		components.ServiceEditDialog(url, s).Render(r.Context(), w)
	}
}
