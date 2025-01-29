package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleResetUserPassword(u auth.UserPasswordResetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		i := auth.UserPasswordResetterInput{Email: r.FormValue("email")}

		err := u.Reset(i)

		if errors.Is(err, auth.ErrEmptyEmail) {
			_http.SendBadRequest(w, "E-mail do utilizador vazio")
			return
		}

		if errors.Is(err, auth.ErrInvalidEmailFormat) {
			_http.SendBadRequest(w, "O e-mail fornecido é inválido")
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
