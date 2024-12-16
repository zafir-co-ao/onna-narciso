package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/web/crm/components"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleUpdateCustomerDialog(f crm.CustomerFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")
		url := r.FormValue("hx-put")

		o, err := f.FindByID(id)

		if errors.Is(err, crm.ErrCustomerNotFound) {
			_http.SendNotFound(w, "Cliente não encontrado")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		_http.SendOk(w)
		components.HandleCustomerUpdateDialog(url, o).Render(r.Context(), w)
	}
}
