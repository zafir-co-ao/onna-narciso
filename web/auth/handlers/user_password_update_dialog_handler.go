package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	"github.com/zafir-co-ao/onna-narciso/web/auth/components"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleUserPasswordUpdateDialog(u auth.UserFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.FormValue("hx-put")
		id := r.FormValue("id")

		o, err := u.FindByID(id)

		if errors.Is(err, auth.ErrUserNotFound) {
			_http.SendNotFound(w, "Utilizador não encontrado")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		_http.SendOk(w)
		components.HandleUpdateUserPasswordDialog(url, o).Render(r.Context(), w)
	}
}
