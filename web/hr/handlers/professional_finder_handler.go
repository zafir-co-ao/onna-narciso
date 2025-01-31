package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	"github.com/zafir-co-ao/onna-narciso/web/auth/handlers"
	"github.com/zafir-co-ao/onna-narciso/internal/hr"
	"github.com/zafir-co-ao/onna-narciso/web/hr/pages"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleFindProfessionals(pf hr.ProfessionalFinder, uf auth.UserFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		o, err := pf.FindAll()

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		au, ok := handlers.GetAuthenticatedUser(w, r, uf)
		if !ok {
			return
		}

		_http.SendOk(w)
		pages.ListProfessionals(o, au).Render(r.Context(), w)
	}
}
