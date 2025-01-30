package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/auth"
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
	uf auth.UserFinder,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := sc.Create(r.FormValue("appointment-id"))

		if errors.Is(err, sessions.ErrInvalidCheckinDate) {
			_http.SendBadRequest(w, fmt.Sprintf("Não é possível fazer o CheckIn nesta data: %s", date.Today().String()))
			return
		}

		if errors.Is(err, sessions.ErrAppointmentCanceled) {
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

		cookie, _ := r.Cookie("userID")
		uid := cookie.Value

		au, err := uf.FindByID(uid)

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		if errors.Is(err, auth.ErrUserNotFound) {
			_http.SendNotFound(w, "Utilizador não encontrado")
			return
		}

		_http.SendOk(w)
		opts := components.CombineAppointmentsWithSessions(appointments, sessions)
		pages.DailyAppointments(date, opts, au).Render(r.Context(), w)
	}
}
