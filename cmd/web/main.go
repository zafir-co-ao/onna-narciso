package main

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/tests/stubs"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/event"
	"github.com/zafir-co-ao/onna-narciso/web"
)

func main() {
	bus := event.NewInmemEventBus()
	repo := inmem.NewAppointmentRepository()
	cacl := stubs.CustomerAclStub{}
	pacl := stubs.Pacl
	sacl := stubs.Sacl

	s := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)
	c := scheduling.NewAppointmentCanceler(repo, bus)
	f := scheduling.NewAppointmentFinder(repo)
	r := scheduling.NewAppointmentRescheduler(repo, bus)

	http.Handle("/", web.NewRouter(s, c, f, r))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
