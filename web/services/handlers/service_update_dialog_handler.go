package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/web/services/components"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleUpdateServiceDialog(u services.ServiceFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.FormValue("hx-put")

		o, err := u.FindByID(r.FormValue("id"))

		if errors.Is(err, services.ErrServiceNotFound) {
			_http.SendNotFound(w, "Serviço não encontrado")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		components.ServiceUpdateDialog(url, o).Render(r.Context(), w)
		_http.SendOk(w)
	}
}
