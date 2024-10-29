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
	f scheduling.AppointmentFinder,
	r scheduling.AppointmentRescheduler,
) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/appointments/{id}", handlers.NewAppointmentFinderHandler(f))
	mux.HandleFunc("/appointments/schedule", handlers.NewAppointmentSchedulerHandler(s))
	mux.HandleFunc("/appointments/reschedule", handlers.NewAppointmentReschedulerHandler(r))
	mux.HandleFunc("/appointments/{id}/cancel", handlers.NewAppointmentCancelerHandler(c))
	mux.HandleFunc("/daily-view", handlers.HandleDailyView)
	mux.HandleFunc("/weekly-appointments", handlers.HandleWeeklyAppointments)

	mux.HandleFunc("/", handlers.NewStaticHandler())

	return mux
}
