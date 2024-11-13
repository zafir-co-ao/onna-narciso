package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/pages"
	"github.com/zafir-co-ao/onna-narciso/web/shared"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleCloseSession(
	sc sessions.Closer,
	sf sessions.Finder,
	df scheduling.DailyAppointmentsFinder,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		input := sessions.CloserInput{SessionID: r.PathValue("id")}

		err := sc.Close(input)

		if errors.Is(sessions.ErrSessionClosed, err) {
			_http.SendBadRequest(w, "A sessão já foi finalizada")
			return
		}

		if errors.Is(sessions.ErrSessionNotFound, err) {
			_http.SendNotFound(w, "A sessão não foi encontrada")
			return
		}

		if errors.Is(sessions.ErrServiceNotFound, err) {
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

		_http.SendOk(w)
		opts := shared.CombineAppointmentsAndSessions(appointments, sessions)
		pages.DailyAppointments(date, opts).Render(r.Context(), w)
	}
}
