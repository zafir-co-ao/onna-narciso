package web

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/web/handlers"
)

var cwd string

func NewRouter() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/weekly-appointments", handlers.HandleWeeklyAppointments)
	mux.HandleFunc("/", handlers.NewStaticHandler())

	return mux
}
