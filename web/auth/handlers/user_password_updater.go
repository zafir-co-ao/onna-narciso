package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleUpdateUserPassword(u auth.UserPasswordUpdater) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		i := auth.UserPasswordUpdaterInput{
			UserID:   r.PathValue("id"),
			Password: r.FormValue("password"),
		}

		err := u.Update(i)

		if errors.Is(err, auth.ErrEmptyPassword) {
			_http.SendBadRequest(w, "Palavra-passe do utilizador vazia")
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

		w.Header().Set("X-Reload-Page", "ReloadPage")
		_http.SendOk(w)
	}
}
