package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/web/crm/pages"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleFindCustomer(cu crm.CustomerFinder, uu auth.UserFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		o, err := cu.FindAll()

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		if errors.Is(err, crm.ErrCustomerNotFound) {
			_http.SendNotFound(w, "Clinte não encontrado")
			return
		}

		cookie, _ := r.Cookie("userID")
		uid := cookie.Value

		au, err := uu.FindByID(uid)

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		if errors.Is(err, auth.ErrUserNotFound) {
			_http.SendNotFound(w, "Utilizador não encontrado")
			return
		}

		_http.SendOk(w)
		pages.ListCustomers(o, au).Render(r.Context(), w)
	}
}
