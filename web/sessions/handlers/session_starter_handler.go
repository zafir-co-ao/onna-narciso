package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions"
	"github.com/zafir-co-ao/onna-narciso/web/auth/handlers"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/pages"
	"github.com/zafir-co-ao/onna-narciso/web/shared/components"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleStartSession(
	ss sessions.SessionStarter,
	sf sessions.SessionFinder,
	dg scheduling.DailyAppointmentsFinder,
	uf auth.UserFinder,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := ss.Start(r.FormValue("session-id"))

		if errors.Is(err, sessions.ErrSessionStarted) {
			_http.SendBadRequest(w, "A sessão já foi iniciada")
			return
		}

		if errors.Is(err, sessions.ErrSessionNotFound) {
			_http.SendNotFound(w, "Sessão não encontrada")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		date := r.FormValue("date")
		if date == "" {
			date = time.Now().Format("2006-01-02")
		}

		appointments, err := dg.Find(date)
		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		appointmentsIDs := xslices.Map(appointments, func(a scheduling.AppointmentOutput) string { return a.ID })

		sessions, err := sf.Find(appointmentsIDs)
		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		au, ok := handlers.GetAuthenticatedUser(w, r, uf)
		if !ok {
			return
		}

		_http.SendOk(w)
		opts := components.CombineAppointmentsWithSessions(appointments, sessions)
		pages.DailyAppointments(date, opts, au).Render(r.Context(), w)
	}
}
