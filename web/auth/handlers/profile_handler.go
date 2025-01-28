package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	"github.com/zafir-co-ao/onna-narciso/web/auth/pages"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleProfilePage(u auth.UserFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		undefinedID := "${id}"

		cookie, _ := r.Cookie("profileID")
		pid := cookie.Value

		if id == undefinedID {
			id = pid
		}

		o, err := u.FindByID(id)

		if errors.Is(err, auth.ErrUserNotFound) {
			_http.SendNotFound(w, "Utilizador não encontrado")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		_http.SendOk(w)
		pages.Profile(o, pid).Render(r.Context(), w)
	}
}
