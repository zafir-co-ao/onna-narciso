package web

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/handlers"
	_sessions "github.com/zafir-co-ao/onna-narciso/web/sessions/handlers"
)

var cwd string

func NewRouter(
	s scheduling.AppointmentScheduler,
	c scheduling.AppointmentCanceler,
	g scheduling.AppointmentGetter,
	r scheduling.AppointmentRescheduler,
	wg scheduling.WeeklyAppointmentsFinder,
	dg scheduling.DailyAppointmentsFinder,
	sc sessions.Creator,
	ss sessions.Starter,
	so sessions.Closer,
	sf sessions.Finder,
) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /appointments", handlers.HandleScheduleAppointment(s))
	mux.HandleFunc("PUT /appointments/{id}", handlers.HandleRescheduleAppointment(r))
	mux.HandleFunc("DELETE /appointments/{id}", handlers.HandleCancelAppointment(c))

	mux.HandleFunc("GET /daily-appointments", handlers.HandleDailyAppointments(dg, sf))
	mux.HandleFunc("GET /weekly-appointments", handlers.HandleWeeklyAppointments(wg))

	mux.HandleFunc("GET /scheduling/dialogs/schedule-appointment-dialog", handlers.HandleScheduleAppointmentDialog())
	mux.HandleFunc("GET /scheduling/dialogs/edit-appointment-dialog/{id}", handlers.HandleEditAppointmentDialog(g))
	mux.HandleFunc("GET /scheduling/daily-appointments-calendar", handlers.HandleDailyAppointmentsCalendar())
	mux.HandleFunc("GET /scheduling/find-professionals/", handlers.HandleFindProfessionals())

	mux.HandleFunc("POST /sessions", _sessions.HandleCreateSession(sc, sf, dg))
	mux.HandleFunc("PUT /sessions/{id}", _sessions.HandleStartSession(ss, sf, dg))
	mux.HandleFunc("DELETE /sessions/{id}", _sessions.HandleCloseSession(so, sf, dg))

	mux.HandleFunc("/", NewStaticHandler())

	return mux
}
