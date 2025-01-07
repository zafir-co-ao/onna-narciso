package handlers

import (
	"errors"
	"net/http"

	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions"
	"github.com/zafir-co-ao/onna-narciso/web/sessions/components"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleCloseSessionDialog(
	ssf sessions.SessionFinder,
	sef services.ServiceFinder,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := ssf.FindByID(r.PathValue("id"))

		if errors.Is(err, sessions.ErrSessionNotFound) {
			_http.SendNotFound(w, "Sessão não encontrada")
			return
		}

		svc, err := sef.FindAll()

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		svc = xslices.Filter(svc, func(service services.ServiceOutput) bool {
			if len(session.Services) == 0 {
				return false
			}

			return service.ID != session.Services[0].ID
		})

		opts := components.SessionCloseOptions{
			Session:  session,
			Services: svc,
			HxDelete: r.FormValue("hx-delete"),
			HxTarget: r.FormValue("hx-target"),
			HxSwap:   r.FormValue("hx-swap"),
		}

		components.SessionCloserDialog(opts).Render(r.Context(), w)
	}
}
