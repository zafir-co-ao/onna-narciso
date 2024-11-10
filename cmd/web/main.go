package main

import (
	"net/http"

	"github.com/kindalus/godx/pkg/event"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/tests/stubs"
	"github.com/zafir-co-ao/onna-narciso/internal/session"
	_session "github.com/zafir-co-ao/onna-narciso/internal/session/adapters/inmem"
	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
	"github.com/zafir-co-ao/onna-narciso/web"
)

func main() {
	bus := event.NewEventBus()
	repo := inmem.NewAppointmentRepository(testdata.Appointments...)
	cacl := stubs.CustomerACLStub{}
	pacl := stubs.Pacl
	sacl := stubs.Sacl

	s := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)
	c := scheduling.NewAppointmentCanceler(repo, bus)
	g := scheduling.NewAppointmentGetter(repo)
	r := scheduling.NewAppointmentRescheduler(repo, bus)
	wg := scheduling.NewWeeklyAppointmentsGetter(repo)

	fs := session.FakeServiceACL{}
	sRepo := _session.NewSessionRepository()
	sc := session.NewSessionCloser(sRepo, fs, bus)

	http.Handle("/", web.NewRouter(s, c, g, r, wg, sc))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
