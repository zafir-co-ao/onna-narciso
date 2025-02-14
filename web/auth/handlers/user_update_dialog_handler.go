package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	"github.com/zafir-co-ao/onna-narciso/web/auth/components"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleUpdateUserDialog(uf auth.UserFinder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")
		url := r.FormValue("hx-put")
		hxTarget := r.FormValue("hx-target")
		hxTriggerEvent := r.FormValue("hx-trigger-event")

		u, err := uf.FindByID(id)

		if errors.Is(err, auth.ErrUserNotFound) {
			_http.SendNotFound(w, "Utilizador não encontrado")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		au, ok := HandleGetAuthenticatedUser(w, r, uf)
		if !ok {
			_http.SendUnauthorized(w)
			return
		}

		p := components.UserUpdateParams{
			Url:            url,
			User:           u,
			AuthUser:       au,
			HxTarget:       hxTarget,
			HxTriggerEvent: hxTriggerEvent,
		}

		_http.SendOk(w)
		components.HandleUpdateUserDialog(p).Render(r.Context(), w)
	}
}
