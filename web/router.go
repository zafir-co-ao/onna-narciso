package web

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/web/handlers"
)

var cwd string

func NewRouter() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/week-view", handlers.HandleWeekView)
	mux.HandleFunc("/", handlers.NewStaticHandler())

	return mux
}
