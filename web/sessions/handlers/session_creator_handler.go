package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/pages"
	"github.com/zafir-co-ao/onna-narciso/web/shared/components"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleCreateSession(
	sc sessions.SessionCreator,
	sf sessions.SessionFinder,
	dg scheduling.DailyAppointmentsFinder,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := sc.Create(r.FormValue("appointment-id"))

		if errors.Is(err, sessions.ErrInvalidCheckinDate) {
			_http.SendBadRequest(w, fmt.Sprintf("Não é possível fazer o CheckIn nesta data: %s", date.Today().String()))
			return
		}

		if errors.Is(sessions.ErrAppointmentCanceled, err) {
			_http.SendBadRequest(w, "A marcação foi cancelada. Não é possível fazer o Check In")
			return
		}

		if errors.Is(err, sessions.ErrAppointmentNotFound) {
			_http.SendNotFound(w, "Marcação não encontrada")
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
		opts := components.CombineAppointmentsWithSessions(appointments, sessions)
		pages.DailyAppointments(date, opts).Render(r.Context(), w)
	}
}
