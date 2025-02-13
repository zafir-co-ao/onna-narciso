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
		hxTarget := r.FormValue("hx-target")
		hxTriggerEvent := r.FormValue("hx-trigger-event")

		o, err := u.FindByID(r.FormValue("id"))

		if errors.Is(err, auth.ErrUserNotFound) {
			_http.SendNotFound(w, "Utilizador não encontrado")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		p := components.UserPasswordUpdateParams{
			Url:            url,
			User:           o,
			HxTarget:       hxTarget,
			HxTriggerEvent: hxTriggerEvent,
		}

		_http.SendOk(w)
		components.HandleUpdateUserPasswordDialog(p).Render(r.Context(), w)
	}
}
