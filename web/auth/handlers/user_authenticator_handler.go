package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleAuthenticateUser(u auth.UserAuthenticator) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		i := auth.UserAuthenticatorInput{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}

		_, err := u.Authenticate(i)

		if errors.Is(err, auth.ErrAuthenticationFailed) {
			_http.SendBadRequest(w, "Credênciais inválidas")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		_http.SendOk(w)
	}
}
