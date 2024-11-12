package main

import (
	"net/http"

	"github.com/kindalus/godx/pkg/event"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/stubs"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/event"
	_session "github.com/zafir-co-ao/onna-narciso/internal/session/adapters/inmem"

	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
	"github.com/zafir-co-ao/onna-narciso/web"
)

func main() {
	bus := event.NewEventBus()
	repo := inmem.NewAppointmentRepository(testdata.Appointments...)
<<<<<<< HEAD
	cacl := stubs.NewCustomersACL()
	pacl := stubs.NewProfessionalsACL()
	sacl := stubs.NewServicesACL()
=======
	cacl := stubs.CustomerACLStub{}
	pacl := stubs.Pacl
	sacl := stubs.Sacl
>>>>>>> 1da440272911f860c6c2354128b8ed904f112aa1

	s := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)
	c := scheduling.NewAppointmentCanceler(repo, bus)
	g := scheduling.NewAppointmentGetter(repo)
	r := scheduling.NewAppointmentRescheduler(repo, bus)
	wg := scheduling.NewWeeklyAppointmentsGetter(repo)
	dg := scheduling.NewDailyAppointmentsGetter(repo)

	fs := session.FakeServiceACL{}
	sRepo := _session.NewSessionRepository(testdata.Sessions...)
	sc := session.NewSessionCreator(sRepo, bus)
	so := session.NewSessionCloser(sRepo, fs, bus)
	sf := session.NewSessionFinder(sRepo)
	ss := session.NewSessionStarter(sRepo, bus)

	http.Handle("/", web.NewRouter(s, c, g, r, wg, dg, sc, ss, so, sf))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
