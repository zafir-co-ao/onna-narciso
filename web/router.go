package web

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/web/handlers"
)

var cwd string

func NewRouter(
	s scheduling.AppointmentScheduler,
	c scheduling.AppointmentCanceler,
	f scheduling.AppointmentGetter,
	r scheduling.AppointmentRescheduler,
	wg scheduling.WeeklyAppointmentsFinder,
) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/appointments/{id}", handlers.NewAppointmentFinderHandler(f))
	mux.HandleFunc("/appointments/schedule", handlers.NewAppointmentSchedulerHandler(s))
	mux.HandleFunc("/appointments/reschedule", handlers.NewAppointmentReschedulerHandler(r))
	mux.HandleFunc("/appointments/{id}/cancel", handlers.NewAppointmentCancelerHandler(c))
	mux.HandleFunc("/daily-view", handlers.HandleDailyView)
	mux.HandleFunc("/weekly-appointments", handlers.HandleWeeklyAppointments(wg))
	mux.HandleFunc("/appointment-form", handlers.HandleAppointmentForm)

	mux.HandleFunc("/", handlers.NewStaticHandler())

	return mux
}
