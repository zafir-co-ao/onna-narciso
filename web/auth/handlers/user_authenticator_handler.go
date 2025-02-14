package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleAuthenticateUser(u auth.UserAuthenticator) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		i := auth.UserAuthenticatorInput{
			Username: strings.ReplaceAll(r.FormValue("username"), " ", ""),
			Password: r.FormValue("password"),
		}

		o, err := u.Authenticate(i)

		if errors.Is(err, auth.ErrAuthenticationFailed) {
			_http.SendUnauthorized(w)
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		cookie := &http.Cookie{
			Name:     "userID",
			Value:    o.ID,
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
		}

		http.SetCookie(w, cookie)

		w.Header().Set("HX-Redirect", "/")
		_http.SendOk(w)
	}
}
