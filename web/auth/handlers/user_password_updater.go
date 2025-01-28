package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleUpdateUserPassword(u auth.UserPasswordUpdater) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		i := auth.UserPasswordUpdaterInput{
			UserID:               id,
			OldPassword:          r.FormValue("old-password"),
			NewPassword:          r.FormValue("new-password"),
			ConfirmationPassword: r.FormValue("confirmation-password"),
		}

		err := u.Update(i)

		if errors.Is(err, auth.ErrInvalidOldPassword) {
			_http.SendBadRequest(w, "Palavra-passe antiga inválida")
			return
		}

		if errors.Is(err, auth.ErrEmptyPassword) {
			_http.SendBadRequest(w, "Palavra-passe do utilizador vazia")
			return
		}

		if errors.Is(err, auth.ErrInvalidConfirmationPassword) {
			_http.SendBadRequest(w, "A palavra-passe de confirmação deve ser a mesma que a nova palavra-passe")
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

		cookie, _ := r.Cookie("userID")
		uid := cookie.Value

		if id == uid {
			UpdateUserPasswordMiddleware(w, r)
			return
		}

		cookie = &http.Cookie{
			Name:     "profileID",
			Value:    id,
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
		}

		http.SetCookie(w, cookie)
		w.Header().Set("X-Reload-Page", "ReloadPage")
		_http.SendOk(w)
	}
}

func UpdateUserPasswordMiddleware(w http.ResponseWriter, r *http.Request) {
	HandleLogoutUser(w, r)
}
