package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleAuthenticateUser(u auth.UserAuthenticator) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		i := auth.UserAuthenticatorInput{
			Username: r.FormValue("username"),
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
			Expires:  time.Now().Add(30 * time.Minute),
			HttpOnly: true,
			Path:     "/",
		}

		http.SetCookie(w, cookie)
		w.Header().Set("HX-Redirect", "/")
		_http.SendOk(w)
	}
}
