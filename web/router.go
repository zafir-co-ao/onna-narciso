package web

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions"
	_auth "github.com/zafir-co-ao/onna-narciso/web/auth/handlers"
	_crm "github.com/zafir-co-ao/onna-narciso/web/crm/handlers"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/handlers"
	_services "github.com/zafir-co-ao/onna-narciso/web/services/handlers"
	_sessions "github.com/zafir-co-ao/onna-narciso/web/sessions/handlers"
)

type UsecasesParams struct {
	AppointmentScheduler     scheduling.AppointmentScheduler
	AppointmentRescheduler   scheduling.AppointmentRescheduler
	AppointmentCanceler      scheduling.AppointmentCanceler
	AppointmentGetter        scheduling.AppointmentGetter
	WeeklyAppointmentsFinder scheduling.WeeklyAppointmentsFinder
	DailyAppointmentsFinder  scheduling.DailyAppointmentsFinder
	SessionCreator           sessions.Creator
	SessionStarter           sessions.Starter
	SessionCloser            sessions.Closer
	SessionFinder            sessions.Finder
	ServiceFinder            services.ServiceFinder
	ServiceCreator           services.ServiceCreator
	ServiceUpdater           services.ServiceUpdater
	ServiceGetter            services.ServiceGetter
	CustomerCreator          crm.CustomerCreator
	CustomerUpdater          crm.CustomerUpdater
	CustomerFinder           crm.CustomerFinder
	CustomerGetter           crm.CustomerGetter
	UserAutheticator         auth.UserAuthenticator
	UserFinder               auth.UserFinder
	UserCreator              auth.UserCreator
}

func NewRouter(u UsecasesParams) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /appointments", handlers.HandleScheduleAppointment(u.AppointmentScheduler))
	mux.HandleFunc("PUT /appointments/{id}", handlers.HandleRescheduleAppointment(u.AppointmentRescheduler))
	mux.HandleFunc("DELETE /appointments/{id}", handlers.HandleCancelAppointment(u.AppointmentCanceler))

	mux.HandleFunc("GET /daily-appointments", handlers.HandleDailyAppointments(u.DailyAppointmentsFinder, u.SessionFinder))
	mux.HandleFunc("GET /weekly-appointments", handlers.HandleWeeklyAppointments(u.WeeklyAppointmentsFinder))

	mux.HandleFunc("GET /scheduling/dialogs/schedule-appointment-dialog", handlers.HandleScheduleAppointmentDialog())
	mux.HandleFunc("GET /scheduling/dialogs/edit-appointment-dialog/{id}", handlers.HandleEditAppointmentDialog(u.AppointmentGetter))
	mux.HandleFunc("GET /scheduling/daily-appointments-calendar", handlers.HandleDailyAppointmentsCalendar())
	mux.HandleFunc("GET /scheduling/find-professionals/", handlers.HandleFindProfessionals())

	mux.HandleFunc("POST /sessions", _sessions.HandleCreateSession(u.SessionCreator, u.SessionFinder, u.DailyAppointmentsFinder))
	mux.HandleFunc("PUT /sessions/{id}", _sessions.HandleStartSession(u.SessionStarter, u.SessionFinder, u.DailyAppointmentsFinder))
	mux.HandleFunc("DELETE /sessions/{id}", _sessions.HandleCloseSession(u.SessionCloser, u.SessionFinder, u.DailyAppointmentsFinder))

	mux.HandleFunc("GET /services", _services.HandleFindServices(u.ServiceFinder))
	mux.HandleFunc("POST /services", _services.HandleCreateService(u.ServiceCreator))
	mux.HandleFunc("PUT /services/{id}", _services.HandleUpdateService(u.ServiceUpdater))
	mux.HandleFunc("GET /services/dialogs/create-service-dialog", _services.HandleCreateServiceDialog)
	mux.HandleFunc("GET /services/dialogs/edit-service-dialog", _services.HandleUpdateServiceDialog(u.ServiceGetter))

	mux.HandleFunc("POST /customers", _crm.HandleCreateCustomer(u.CustomerCreator))
	mux.HandleFunc("GET /customers", _crm.HandleFindCustomer(u.CustomerFinder))
	mux.HandleFunc("PUT /customers/{id}", _crm.HandleUpdateCustomer(u.CustomerUpdater))
	mux.HandleFunc("GET /customers/dialogs/create-customer-dialog", _crm.HandleCreateCustomerDialog)
	mux.HandleFunc("GET /customers/dialogs/edit-customer-dialog", _crm.HandleUpdateCustomerDialog(u.CustomerGetter))

	mux.HandleFunc("GET /auth/login", _auth.HandleLoginPage)
	mux.HandleFunc("GET /auth/logout", _auth.HandleLogoutUser)
	mux.HandleFunc("POST /auth/login", _auth.HandleAuthenticateUser(u.UserAutheticator))
	mux.HandleFunc("GET /auth/users", _auth.HandleFindUsers(u.UserFinder))
	mux.HandleFunc("POST /auth/users", _auth.HandleCreateUser(u.UserCreator))
	mux.HandleFunc("GET /users/dialogs/create-user-dialog", _auth.HandleUserCreateDialog)

	mux.HandleFunc("/", NewStaticHandler())

	return mux
}

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("username")
		path := r.URL.Path

		if err != nil && path == "/auth/login" {
			w.Header().Set("HX-Redirect", "/auth/login")
			next.ServeHTTP(w, r)
			return
		}

		if err != nil && path != "/auth/login" {
			w.Header().Set("HX-Redirect", "/auth/login")
			next.ServeHTTP(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
