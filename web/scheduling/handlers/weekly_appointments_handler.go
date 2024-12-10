package handlers

import (
	"errors"
	"net/http"
	"slices"
	"time"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
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

	//TODO - Quando tiver o repositório de profissionais, permitir apenas serviços que o profissional atende
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

func HandleWeeklyAppointments(g scheduling.WeeklyAppointmentsFinder) func(w http.ResponseWriter, r *http.Request) {

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

		professionals := make([]scheduling.Professional, 0)
		if serviceID != "all" {
			professionals = testdata.Professionals
			//TODO - Utilizar o repositório de profissionais para filtrar os profissionais que atendem o serviço
			tmp := make([]scheduling.Professional, 0)
			for _, professional := range professionals {
				if slices.Contains(professional.ServicesIDS, nanoid.ID(serviceID)) {
					tmp = append(tmp, professional)
				}
			}
			professionals = tmp
		}

		appointments, err := findApppointments(g, date, serviceID, professionalID)
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
			Services:       testdata.Services,
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
