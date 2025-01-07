package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kindalus/godx/pkg/event"
	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/hr"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions"

	_hr_stubs "github.com/zafir-co-ao/onna-narciso/internal/hr/stubs"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/stubs"
	_stubs "github.com/zafir-co-ao/onna-narciso/internal/sessions/stubs"

	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
	"github.com/zafir-co-ao/onna-narciso/web"
)

func main() {
	bus := event.NewEventBus()

	bus.SubscribeFunc(scheduling.EventAppointmentScheduled, sendNotification)

	cacl := stubs.NewCRMServiceACL()
	pacl := stubs.NewHRServiceACL()
	sacl := stubs.NewServicesServiceACL()
	aacl := _stubs.NewSchedulingServiceACL()
	ssacl := _stubs.NewServicesServiceACL()
	hrsacl := _hr_stubs.NewServicesServiceACL()

	appointmentRepo := scheduling.NewAppointmentRepository(testdata.Appointments...)
	sessionRepo := sessions.NewInmemRepository(testdata.Sessions...)
	serviceRepo := services.NewInmemRepository(testdata.ServicesDummies...)
	customerRepo := crm.NewInmemRepository(testdata.CustomersDummies...)
	userRepo := auth.NewInmemRepository(testdata.Users...)
	professionalRepo := hr.NewInmemProfessionalRepository(testdata.ProfessionalsDummies...)

	u := web.UsecasesParams{
		AppointmentScheduler:     scheduling.NewAppointmentScheduler(appointmentRepo, cacl, pacl, sacl, bus),
		AppointmentRescheduler:   scheduling.NewAppointmentRescheduler(appointmentRepo, pacl, sacl, bus),
		AppointmentCanceler:      scheduling.NewAppointmentCanceler(appointmentRepo, bus),
		AppointmentGetter:        scheduling.NewAppointmentFinder(appointmentRepo),
		WeeklyAppointmentsFinder: scheduling.NewWeeklyAppointmentsFinder(appointmentRepo),
		DailyAppointmentsFinder:  scheduling.NewDailyAppointmentsFinder(appointmentRepo),
		SessionCreator:           sessions.NewSessionCreator(sessionRepo, bus, aacl),
		SessionStarter:           sessions.NewSessionStarter(sessionRepo, bus),
		SessionCloser:            sessions.NewSessionCloser(sessionRepo, ssacl, bus),
		SessionFinder:            sessions.NewSessionFinder(sessionRepo),
		ServiceCreator:           services.NewServiceCreator(serviceRepo, bus),
		ServiceFinder:            services.NewServiceFinder(serviceRepo),
		ServiceUpdater:           services.NewServiceUpdater(serviceRepo, bus),
		CustomerCreator:          crm.NewCustomerCreator(customerRepo, bus),
		CustomerUpdater:          crm.NewCustomerUpdater(customerRepo, bus),
		CustomerFinder:           crm.NewCustomerFinder(customerRepo),
		UserAutheticator:         auth.NewUserAuthenticator(userRepo),
		UserFinder:               auth.NewUserFinder(userRepo),
		UserCreator:              auth.NewUserCreator(userRepo, bus),
		ProfessionalCreator:      hr.NewProfessionalCreator(professionalRepo, hrsacl, bus),
		ProfessionalFinder:       hr.NewProfessionalFinder(professionalRepo),
		ProfessionalUpdater:      hr.NewProfessionalUpdater(professionalRepo, hrsacl, bus),
	}

	r := web.NewRouter(u)
	http.Handle("/", web.AuthenticationMiddleware(r))

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
