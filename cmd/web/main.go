package main

import (
	"net/http"

	"github.com/kindalus/godx/pkg/event"
	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/stubs"
	_stubs "github.com/zafir-co-ao/onna-narciso/internal/sessions/stubs"

	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
	"github.com/zafir-co-ao/onna-narciso/web"
)

func main() {
	bus := event.NewEventBus()
	cacl := stubs.NewCRMServiceACL()
	pacl := stubs.NewHRServiceACL()
	sacl := stubs.NewServicesServiceACL()
	aacl := _stubs.NewSchedulingServiceACL()
	ssacl := _stubs.NewServicesServiceACL()

	appointmentRepo := scheduling.NewAppointmentRepository(testdata.Appointments...)
	sessionRepo := sessions.NewInmemRepository(testdata.Sessions...)
	serviceRepo := services.NewInmemRepository(testdata.ServicesDummies...)
	customerRepo := crm.NewInmemRepository(testdata.CustomersDummies...)
	userRepo := auth.NewInmemRepository(testdata.Users...)

	u := web.UsecasesParams{
		AppointmentScheduler:     scheduling.NewAppointmentScheduler(appointmentRepo, cacl, pacl, sacl, bus),
		AppointmentRescheduler:   scheduling.NewAppointmentRescheduler(appointmentRepo, pacl, sacl, bus),
		AppointmentCanceler:      scheduling.NewAppointmentCanceler(appointmentRepo, bus),
		AppointmentFinder:        scheduling.NewAppointmentFinder(appointmentRepo),
		WeeklyAppointmentsFinder: scheduling.NewWeeklyAppointmentsFinder(appointmentRepo),
		DailyAppointmentsFinder:  scheduling.NewDailyAppointmentsFinder(appointmentRepo),
		SessionCreator:           sessions.NewSessionCreator(sessionRepo, aacl, bus),
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
		UserUpdater:              auth.NewUserUpdater(userRepo, bus),
		UserPasswordUpdater:      auth.NewUserPasswordUpdater(userRepo, bus),
	}

	r := web.NewRouter(u)
	http.Handle("/", web.AuthenticationMiddleware(r))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
