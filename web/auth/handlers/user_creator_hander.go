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
			UserID:   uid,
			Username: strings.TrimSpace(r.FormValue("username")),
			Password: r.FormValue("password"),
			Role:     r.FormValue("role"),
		}

		_, err := u.Create(i)

		if errors.Is(err, auth.ErrEmptyUsername) {
			_http.SendBadRequest(w, "Nome do utilizador vazio")
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

		if errors.Is(err, auth.ErrUserNotFound) {
			_http.SendNotFound(w, "Utilizador não encontrado")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		w.Header().Set("X-Reload-Page", "ReloadPage")
		_http.SendCreated(w)
	}
}
