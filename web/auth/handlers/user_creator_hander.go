package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleCreateUser(u auth.UserCreator) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("userID")
		uid := cookie.Value

		i := auth.UserCreatorInput{
			UserID:      uid,
			Username:    strings.ReplaceAll(r.FormValue("username"), " ", ""),
			Email:       r.FormValue("email"),
			PhoneNumber: r.FormValue("phone-number"),
			Password:    r.FormValue("password"),
			Role:        r.FormValue("role"),
		}

		_, err := u.Create(i)

		if errors.Is(err, auth.ErrEmptyUsername) {
			_http.SendBadRequest(w, "Nome do utilizador vazio")
			return
		}

		if errors.Is(err, auth.ErrEmptyEmail) {
			_http.SendBadRequest(w, "E-mail do utilizador vazio")
			return
		}

		if errors.Is(err, auth.ErrInvalidEmailFormat) {
			_http.SendBadRequest(w, "O e-mail fornecido é inválido")
			return
		}

		if errors.Is(err, auth.ErrEmptyPhoneNumber) {
			_http.SendBadRequest(w, "Telefone do utilizador vazio")
			return
		}

		if errors.Is(err, auth.ErrEmptyPassword) {
			_http.SendBadRequest(w, "Palavra-passe do utilizador vazia")
			return
		}

		if errors.Is(err, auth.ErrRoleNotAllowed) {
			_http.SendBadRequest(w, "Perfil de utilizador não permitido")
			return
		}

		if errors.Is(err, auth.ErrUserNotAllowed) {
			_http.SendBadRequest(w, "Não pode criar novos utilizadores")
			return
		}

		if errors.Is(err, auth.ErrOnlyUniqueUsername) {
			_http.SendBadRequest(w, "O nome do utilizador deve ser único")
			return
		}

		if errors.Is(err, auth.ErrOnlyUniqueEmail) {
			_http.SendBadRequest(w, "O e-mail do utilizador deve ser único")
			return
		}

		if errors.Is(err, auth.ErrOnlyUniquePhoneNumber) {
			_http.SendBadRequest(w, "O número de telefone do utilizador deve ser único")
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

		_http.SendCreated(w)
	}
}
