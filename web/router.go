package web

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/session"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/handlers"
	_session "github.com/zafir-co-ao/onna-narciso/web/session/handlers"
)

var cwd string

func NewRouter(
	s scheduling.AppointmentScheduler,
	c scheduling.AppointmentCanceler,
	g scheduling.AppointmentGetter,
	r scheduling.AppointmentRescheduler,
	wg scheduling.WeeklyAppointmentsFinder,
	dg scheduling.DailyAppointmentsFinder,
	sc session.Closer,
	sf session.Finder,
) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /appointments", handlers.HandleScheduleAppointment(s, wg))
	mux.HandleFunc("PUT /appointments/{id}", handlers.HandleRescheduleAppointment(r, wg))
	mux.HandleFunc("DELETE /appointments/{id}", handlers.HandleCancelAppointment(c, wg))

	mux.HandleFunc("GET /daily-appointments", handlers.HandleDailyAppointments(dg, sf))
	mux.HandleFunc("GET /weekly-appointments", handlers.HandleWeeklyAppointments(wg))

	mux.HandleFunc("GET /scheduling/dialogs/schedule-appointment-dialog", handlers.HandleScheduleAppointmentDialog())
	mux.HandleFunc("GET /scheduling/dialogs/edit-appointment-dialog/{id}", handlers.HandleEditAppointmentDialog(g))

	mux.HandleFunc("DELETE /sessions/{id}", _session.HandleCloseSession(sc, sf, dg))

	mux.HandleFunc("/", NewStaticHandler())

	return mux
}
