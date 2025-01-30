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
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/pages"
	"github.com/zafir-co-ao/onna-narciso/web/shared/components"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleCloseSession(
	sc sessions.SessionCloser,
	sf sessions.SessionFinder,
	df scheduling.DailyAppointmentsFinder,
	uf auth.UserFinder,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		i := sessions.SessionCloserInput{
			SessionID: r.PathValue("id"),
			Services:  toSessionCloserServicesInput(r),
			Gift:      r.FormValue("gift"),
		}

		err = sc.Close(i)

		if errors.Is(err, sessions.ErrSessionClosed) {
			_http.SendBadRequest(w, "A sessão já foi finalizada")
			return
		}

		if errors.Is(err, sessions.ErrSessionNotFound) {
			_http.SendNotFound(w, "A sessão não foi encontrada")
			return
		}

		if errors.Is(err, sessions.ErrServiceNotFound) {
			_http.SendNotFound(w, "Serviço não encontrado")
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

		appointments, err := df.Find(date)

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

func toSessionCloserServicesInput(r *http.Request) []sessions.SessionCloserServiceInput {
	var services []sessions.SessionCloserServiceInput

	for _, s := range r.Form["service-id"] {
		s := sessions.SessionCloserServiceInput{
			ServiceID: s,
			Discount:  r.FormValue(fmt.Sprintf("service-discount-%v", s)),
		}

		services = append(services, s)
	}
	return services
}
