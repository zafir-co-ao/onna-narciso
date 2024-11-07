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

	mux.HandleFunc("/appointments/{id}", handlers.NewAppointmentGetterHandler(g))
	mux.HandleFunc("/appointments", handlers.NewAppointmentSchedulerHandler(s))
	mux.HandleFunc("/appointments/reschedule", handlers.NewAppointmentReschedulerHandler(r))
	mux.HandleFunc("/appointments/{id}/cancel", handlers.NewAppointmentCancelerHandler(c))
	mux.HandleFunc("/daily-appointments", handlers.HandleDailyAppointments)
	mux.HandleFunc("/weekly-appointments", handlers.HandleWeeklyAppointments(wg))
	mux.HandleFunc("/appointment-dialog", handlers.HandleAppointmentDialog())

	mux.HandleFunc("/", NewStaticHandler())

	return mux
}
