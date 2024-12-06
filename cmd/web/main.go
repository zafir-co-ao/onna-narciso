package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kindalus/godx/pkg/event"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions"

	_crm "github.com/zafir-co-ao/onna-narciso/internal/crm/adapters/inmem"
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

	cacl := stubs.NewCustomersACL()
	pacl := stubs.NewProfessionalsACL()
	sacl := stubs.NewServicesACL()
	aacl := _stubs.NewAppointmentsACL()
	ssacl := _stubs.NewServicesACL()

	appointmentRepo := inmem.NewAppointmentRepository(testdata.Appointments...)
	sessionRepo := _sessions.NewSessionRepository(testdata.Sessions...)
	serviceRepo := _services.NewServiceRepository()
	customerRepo := _crm.NewCustomerRepository()

	u := web.UsecasesParams{
		AppointmentScheduler:     scheduling.NewAppointmentScheduler(appointmentRepo, cacl, pacl, sacl, bus),
		AppointmentRescheduler:   scheduling.NewAppointmentRescheduler(appointmentRepo, pacl, sacl, bus),
		AppointmentCanceler:      scheduling.NewAppointmentCanceler(appointmentRepo, bus),
		AppointmentGetter:        scheduling.NewAppointmentGetter(appointmentRepo),
		WeeklyAppointmentsFinder: scheduling.NewWeeklyAppointmentsFinder(appointmentRepo),
		DailyAppointmentsFinder:  scheduling.NewDailyAppointmentsFinder(appointmentRepo),
		SessionCreator:           sessions.NewSessionCreator(sessionRepo, bus, aacl),
		SessionStarter:           sessions.NewSessionStarter(sessionRepo, bus),
		SessionCloser:            sessions.NewSessionCloser(sessionRepo, ssacl, bus),
		SessionFinder:            sessions.NewSessionFinder(sessionRepo),
		ServiceCreator:           services.NewServiceCreator(serviceRepo, bus),
		ServiceFinder:            services.NewServiceFinder(serviceRepo),
		CustomerCreator:          crm.NewCustomerCreator(customerRepo, bus),
	}

	http.Handle("/", web.NewRouter(u))

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
