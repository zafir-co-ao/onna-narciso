package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	"github.com/zafir-co-ao/onna-narciso/web/auth/pages"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleFindUsers(uf auth.UserFinder, ug auth.UserGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("userID")

		uid := cookie.Value
		o, err := uf.Find(uid)

		// if errors.Is(err, auth.ErrUserNotAllowed) {
		// 	_http.SendBadRequest(w, "Acesso negado")
		// 	return
		// }

		if errors.Is(err, auth.ErrUserNotFound) {
			_http.SendNotFound(w, "Utilizador não encontrado")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		au, _ := ug.Get(uid)

		pages.ListUsers(o, au).Render(r.Context(), w)
	}
}
