package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions"
	_auth "github.com/zafir-co-ao/onna-narciso/web/auth/handlers"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/pages"
	"github.com/zafir-co-ao/onna-narciso/web/shared/components"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleDailyAppointments(
	df scheduling.DailyAppointmentsFinder,
	sf sessions.SessionFinder,
	uf auth.UserFinder,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		date := r.FormValue("date")

		if date == "" {
			date = time.Now().Format("2006-01-02")
		}

		a, err := df.Find(date)
		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		appointmentsIDs := xslices.Map(a, func(a scheduling.AppointmentOutput) string { return a.ID })

		s, err := sf.Find(appointmentsIDs)
		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		au, ok := _auth.HandleGetAuthenticatedUser(w, r, uf)
		if !ok {
			_http.SendUnauthorized(w)
			return
		}

		_http.SendOk(w)
		opts := components.CombineAppointmentsWithSessions(a, s)
		pages.DailyAppointments(date, opts, au).Render(r.Context(), w)
	}
}
