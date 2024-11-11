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

func HandleCreateSession(
	sc session.Creator,
	sf session.Finder,
	dg scheduling.DailyAppointmentsFinder,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := sc.Create(r.FormValue("appointment-id"))

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
		opts := shared.CombineAppointmentsAndSessions(appointments, sessions)
		pages.DailyAppointments(opts).Render(r.Context(), w)
	}
}
