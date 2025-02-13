package handlers

import (
	"errors"
	"net/http"
	"slices"
	"time"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/hr"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/duration"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/pages"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func weeklyAppointmentsServiceChanged(date string, serviceID string) (string, string, string) {

	if serviceID == "all" {
		return date, "all", "all"
	}

	return date, serviceID, "all"
}

func weeklyAppointmentsProfessionalChanged(date, serviceID, professionalID string) (string, string, string) {
	if serviceID == "all" {
		return date, "all", "all"
	}

	if professionalID == "all" {
		return date, serviceID, "all"
	}

	return date, serviceID, professionalID
}

func nextWeek(date string) string {
	d, _ := time.Parse("2006-01-02", date)
	d = d.AddDate(0, 0, 7)
	return d.Format("2006-01-02")
}

func previousWeek(date string) string {
	d, _ := time.Parse("2006-01-02", date)
	d = d.AddDate(0, 0, -7)
	return d.Format("2006-01-02")
}

func HandleWeeklyAppointments(af scheduling.WeeklyAppointmentsFinder, pf hr.ProfessionalFinder, sf services.ServiceFinder) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		date := r.FormValue("date")
		previousDate := r.FormValue("previous-date")
		serviceID := r.FormValue("service-id")
		previousServiceID := r.FormValue("previous-service-id")
		professionalID := r.FormValue("professional-id")
		previousProfessionalID := r.FormValue("previous-professional-id")
		operation := r.FormValue("operation")

		if date == "" {
			date = time.Now().Format("2006-01-02")
		}

		if serviceID == "" {
			serviceID = "all"
			previousServiceID = "all"
		}

		if professionalID == "" {
			professionalID = "all"
			previousProfessionalID = "all"
		}

		if professionalID != previousProfessionalID {
			date, serviceID, professionalID = weeklyAppointmentsProfessionalChanged(date, serviceID, professionalID)
		}

		if serviceID != previousServiceID {
			date, serviceID, professionalID = weeklyAppointmentsServiceChanged(date, serviceID)
		}

		if date != previousDate {
			serviceID, professionalID = "all", "all"
		}

		if operation == "previous-week" {
			date = previousWeek(date)
		}

		if operation == "next-week" {
			date = nextWeek(date)
		}

		professionalOutput, err := pf.FindAll()

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		serviceIDs := make([]nanoid.ID, 0)
		for _, p := range professionalOutput {
			for _, s := range p.Services {
				serviceIDs = append(serviceIDs, nanoid.ID(s.ID))
			}
		}

		professionals := xslices.Map(professionalOutput, func(p hr.ProfessionalOutput) scheduling.Professional {
			return scheduling.Professional{
				ID:          nanoid.ID(p.ID),
				Name:        name.Name(p.Name),
				ServicesIDs: serviceIDs,
			}
		})

		servicesOutput, err := sf.FindAll()

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		services := xslices.Map(servicesOutput, func(s services.ServiceOutput) scheduling.Service {
			return scheduling.Service{
				ID:       nanoid.ID(s.ID),
				Name:     name.Name(s.Name),
				Duration: duration.Duration(s.Duration),
			}
		})


		if serviceID != "all" {
			tmp := make([]scheduling.Professional, 0)
			for _, p := range professionals {
				if slices.Contains(p.ServicesIDs, nanoid.ID(serviceID)) {
					tmp = append(tmp, p)
				}
			}
			professionals = tmp
		}

		appointments, err := findApppointments(af, date, serviceID, professionalID)
		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		opts := pages.WeeklyAppointmentsOptions{
			Date:           date,
			ServiceID:      serviceID,
			ProfessionalID: professionalID,
			StartHour:      6,
			EndHour:        20,
			Days:           5,
			Services:       services,
			Professionals:  professionals,
			Appointments:   appointments,
		}

		if serviceID == "all" {
			professionalID = "all"
		}

		pages.WeeklyAppointments(opts).Render(r.Context(), w)
	}
}

func findApppointments(f scheduling.WeeklyAppointmentsFinder,
	date, serviceID, professionalID string) ([]scheduling.AppointmentOutput, error) {

	professionalIDS := make([]string, 0)
	if professionalID != "all" {
		professionalIDS = append(professionalIDS, professionalID)
	}

	if serviceID == "all" {
		return make([]scheduling.AppointmentOutput, 0), nil
	}

	return f.Find(date, serviceID, professionalIDS)
}
