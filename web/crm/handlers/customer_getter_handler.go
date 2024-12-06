package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleGetCustomer(u crm.CustomerGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")

		_, err := u.Get(id)

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		if errors.Is(err, crm.ErrCustomerNotFound) {
			_http.SendNotFound(w, "Cliente não encontrado")
			return
		}

		_http.SendOk(w)
	}
}
