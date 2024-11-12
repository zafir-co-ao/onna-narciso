package main

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/stubs"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/event"
	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
	"github.com/zafir-co-ao/onna-narciso/web"
)

func main() {
	bus := event.NewInmemEventBus()
	repo := inmem.NewAppointmentRepository(testdata.Appointments...)
	cacl := stubs.NewCustomersACL()
	pacl := stubs.NewProfessionalsACL()
	sacl := stubs.NewServicesACL()

	s := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)
	c := scheduling.NewAppointmentCanceler(repo, bus)
	f := scheduling.NewAppointmentGetter(repo)
	r := scheduling.NewAppointmentRescheduler(repo, bus)
	wg := scheduling.NewWeeklyAppointmentsGetter(repo)

	http.Handle("/", web.NewRouter(s, c, f, r, wg))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
