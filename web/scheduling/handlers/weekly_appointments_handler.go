package handlers

import (
	"net/http"
	"slices"
	"time"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/id"
	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/pages"
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

func HandleWeeklyAppointments(g scheduling.WeeklyAppointmentsFinder) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {

		date := r.FormValue("date")
		previousDate := r.FormValue("previous-date")
		serviceID := r.FormValue("service-id")
		previousServiceID := r.FormValue("previous-service-id")
		professionalID := r.FormValue("professional-id")
		previousProfessionalID := r.FormValue("previous-professional-id")

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
			date, serviceID, professionalID = date, "all", "all"
		}

		professionals := make([]scheduling.Professional, 0)
		if serviceID != "all" {
			professionals = testdata.Professionals
			//TODO - Utilizar o repositório de profissionais para filtrar os profissionais que atendem o serviço
			tmp := make([]scheduling.Professional, 0)
			for _, professional := range professionals {
				if slices.Contains(professional.ServicesIDS, id.ID(serviceID)) {
					tmp = append(tmp, professional)
				}
			}
			professionals = tmp
		}

		appointments, err := findApppointments(g, date, serviceID, professionalID)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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