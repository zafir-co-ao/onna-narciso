package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/web/services/pages"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleFindServices(sf services.ServiceFinder, uf auth.UserFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		o, err := sf.FindAll()

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		cookie, _ := r.Cookie("userID")
		uid := cookie.Value

		au, err := uf.FindByID(uid)

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		if errors.Is(err, auth.ErrUserNotFound) {
			_http.SendNotFound(w, "Utilizador não encontrado")
			return
		}

		_http.SendOk(w)
		pages.ListServices(o, au).Render(r.Context(), w)
	}
}
