package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/session"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/pages"
	"github.com/zafir-co-ao/onna-narciso/web/shared"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleStartSession(
	ss session.Starter,
	sf session.Finder,
	dg scheduling.DailyAppointmentsFinder,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := ss.Start(r.FormValue("session-id"))

		if errors.Is(session.ErrSessionStarted, err) {
			_http.SendBadRequest(w, "A sessão já foi iniciada")
			return
		}

		if errors.Is(session.ErrSessionNotFound, err) {
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

		_http.SendOk(w)
		result := shared.CombineAppointmentsAndSessions(appointments, sessions)
		pages.DailyAppointments(result).Render(r.Context(), w)
	}
}
