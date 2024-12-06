package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/web/crm/pages"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleFindCustomer(u crm.CustomerFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		o, err := u.Find()

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		_http.SendOk(w)
		pages.ListCustomers(o).Render(r.Context(), w)
	}
}
