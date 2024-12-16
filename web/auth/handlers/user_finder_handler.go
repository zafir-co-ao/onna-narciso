package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	"github.com/zafir-co-ao/onna-narciso/web/auth/pages"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleFindUsers(u auth.UserFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("username")

		uid := cookie.Value
		o, err := u.Find(uid)

		if errors.Is(err, auth.ErrUserNotAllowed) {
			_http.SendUnauthorized(w)
			return
		}

		if errors.Is(err, auth.ErrUserNotFound) {
			_http.SendNotFound(w, "Utilizador não encontrado")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		_http.SendOk(w)
		pages.ListUsers(o).Render(r.Context(), w)
	}
}
