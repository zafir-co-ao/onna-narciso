package web

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/handlers"
)

var cwd string

func NewRouter(
	s scheduling.AppointmentScheduler,
	c scheduling.AppointmentCanceler,
	g scheduling.AppointmentGetter,
	r scheduling.AppointmentRescheduler,
	wg scheduling.WeeklyAppointmentsFinder,
) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /appointments", handlers.HandleScheduleAppointment(s))
	mux.HandleFunc("PUT /appointments/{id}", handlers.HandleRescheduleNewAppointment(r))
	mux.HandleFunc("DELETE /appointments/{id}", handlers.HandleCancelAppointment(c))

	mux.HandleFunc("GET /daily-appointments", handlers.HandleDailyAppointments)
	mux.HandleFunc("GET /weekly-appointments", handlers.HandleWeeklyAppointments(wg))

	mux.HandleFunc("GET /scheduling/dialogs/schedule-appointment-dialog", handlers.HandleScheduleAppointmentDialog())
	mux.HandleFunc("GET /scheduling/dialogs/edit-appointment-dialog/{id}", handlers.HandleEditAppointmentDialog(g))

	mux.HandleFunc("/", NewStaticHandler())

	return mux
}
