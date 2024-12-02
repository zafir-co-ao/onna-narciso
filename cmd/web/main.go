package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kindalus/godx/pkg/event"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/stubs"
	_services "github.com/zafir-co-ao/onna-narciso/internal/services/adapters/inmem"
	_sessions "github.com/zafir-co-ao/onna-narciso/internal/sessions/adapters/inmem"
	_stubs "github.com/zafir-co-ao/onna-narciso/internal/sessions/stubs"

	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
	"github.com/zafir-co-ao/onna-narciso/web"
)

func main() {
	bus := event.NewEventBus()

	bus.SubscribeFunc(scheduling.EventAppointmentScheduled, sendNotification)

	repo := inmem.NewAppointmentRepository(testdata.Appointments...)
	cacl := stubs.NewCustomersACL()
	pacl := stubs.NewProfessionalsACL()
	sacl := stubs.NewServicesACL()
	aacl := _stubs.NewAppointmentsACL()

	s := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)
	c := scheduling.NewAppointmentCanceler(repo, bus)
	g := scheduling.NewAppointmentGetter(repo)
	r := scheduling.NewAppointmentRescheduler(repo, pacl, sacl, bus)
	wf := scheduling.NewWeeklyAppointmentsFinder(repo)
	df := scheduling.NewDailyAppointmentsFinder(repo)

	fs := _stubs.NewServicesACL()
	sRepo := _sessions.NewSessionRepository(testdata.Sessions...)
	sc := sessions.NewSessionCreator(sRepo, bus, aacl)
	so := sessions.NewSessionCloser(sRepo, fs, bus)
	sf := sessions.NewSessionFinder(sRepo)
	ss := sessions.NewSessionStarter(sRepo, bus)

	scrRepo := _services.NewServiceRepository()
	scr := services.NewServiceCreator(scrRepo, bus)

	http.Handle("/", web.NewRouter(s, c, g, r, wf, df, sc, ss, so, sf, scr))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func sendNotification(event.Event) {
	// Find your Account SID and Auth Token at twilio.com/console
	// and set the environment variables. See http://twil.io/secure
	// Make sure TWILIO_ACCOUNT_SID and TWILIO_AUTH_TOKEN exists in your environment
	client := twilio.NewRestClient()

	params := &api.CreateMessageParams{}
	params.SetBody("This is the ship that made the Kessel Run in fourteen parsecs?")
	params.SetFrom("+15017122661")
	params.SetTo("+244923641819")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else {
		if resp.Body != nil {
			fmt.Println(*resp.Body)
		} else {
			fmt.Println(resp.Body)
		}
	}
}
