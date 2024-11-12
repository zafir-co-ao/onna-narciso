package main

import (
	"net/http"

	"github.com/kindalus/godx/pkg/event"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/stubs"
	_sessions "github.com/zafir-co-ao/onna-narciso/internal/sessions/adapters/inmem"

	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
	"github.com/zafir-co-ao/onna-narciso/web"
)

func main() {
	bus := event.NewEventBus()
	repo := inmem.NewAppointmentRepository(testdata.Appointments...)
	cacl := stubs.NewCustomersACL()
	pacl := stubs.NewProfessionalsACL()
	sacl := stubs.NewServicesACL()

	s := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)
	c := scheduling.NewAppointmentCanceler(repo, bus)
	g := scheduling.NewAppointmentGetter(repo)
	r := scheduling.NewAppointmentRescheduler(repo, bus)
	wg := scheduling.NewWeeklyAppointmentsGetter(repo)
	dg := scheduling.NewDailyAppointmentsGetter(repo)

	fs := sessions.FakeServiceACL{}
	sRepo := _sessions.NewSessionRepository(testdata.Sessions...)
	sc := sessions.NewSessionCreator(sRepo, bus)
	so := sessions.NewSessionCloser(sRepo, fs, bus)
	sf := sessions.NewSessionFinder(sRepo)
	ss := sessions.NewSessionStarter(sRepo, bus)

	http.Handle("/", web.NewRouter(s, c, g, r, wg, dg, sc, ss, so, sf))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
