package web

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/web/handlers"
)

var cwd string

func NewRouter(s scheduling.AppointmentScheduler, c scheduling.AppointmentCanceler) *http.ServeMux {

	mux := http.NewServeMux()
	mux.HandleFunc("/appointments/schedule", handlers.NewAppointmentSchedulerHandler(s))
	mux.HandleFunc("appointments/{id}/cancel", handlers.NewAppointmentCancelerHandler(c))
	mux.HandleFunc("/week-view", handlers.HandleWeekView)
	mux.HandleFunc("/", handlers.NewStaticHandler())

	return mux
}
