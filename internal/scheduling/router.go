package scheduling

import "net/http"

func NewRouter(s AppointmentScheduler, c AppointmentCanceler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/schedule", NewAppointmentSchedulerHandler(s))
	mux.HandleFunc("/{id}/cancel", NewAppointmentCancelerHandler(c))

	return mux
}
