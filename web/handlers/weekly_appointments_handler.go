package handlers

import (
	"net/http"
	"slices"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/id"
	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
	"github.com/zafir-co-ao/onna-narciso/web/components"
)

var appointments = []scheduling.AppointmentOutput{
	{
		ID:           "1",
		CustomerID:   "1",
		CustomerName: "Paola Oliveira",

		ProfessionalID:   testdata.Professionals[0].ID.String(),
		ProfessionalName: testdata.Professionals[0].Name.String(),
		ServiceID:        testdata.Services[0].ID.String(),
		ServiceName:      testdata.Services[0].Name.String(),

		Date:     "2024-10-10",
		Hour:     "08:00",
		Duration: 180,
	},
	{
		ID:           "2",
		CustomerID:   "2",
		CustomerName: "Juliana Paes",

		ProfessionalID:   testdata.Professionals[1].ID.String(),
		ProfessionalName: testdata.Professionals[1].Name.String(),
		ServiceID:        testdata.Services[2].ID.String(),
		ServiceName:      testdata.Services[2].Name.String(),

		Date:     "2024-10-11",
		Hour:     "10:30",
		Duration: 90,
	},
	{
		ID:           "3",
		CustomerID:   "3",
		CustomerName: "Gisele Bündchen",

		ProfessionalID:   testdata.Professionals[2].ID.String(),
		ProfessionalName: testdata.Professionals[2].Name.String(),
		ServiceID:        testdata.Services[3].ID.String(),
		ServiceName:      testdata.Services[3].Name.String(),

		Date:     "2024-10-10",
		Hour:     "12:00",
		Duration: 60,
	},
}

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

func HandleWeeklyAppointments(w http.ResponseWriter, r *http.Request) {

	date := r.URL.Query().Get("date")
	previousDate := r.URL.Query().Get("previous-date")
	serviceID := r.URL.Query().Get("service-id")
	previousServiceID := r.URL.Query().Get("previous-servic-id")
	professionalID := r.URL.Query().Get("professional-id")
	previousProfessionalID := r.URL.Query().Get("previous-professional-id")

	if date == "" {
		//TODO - Utilizar a data atual
		date = "2024-10-10"
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

	professionals := testdata.Professionals
	if serviceID != "all" {
		//TODO - Utilizar o repositório de profissionais para filtrar os profissionais que atendem o serviço
		tmp := make([]scheduling.Professional, 0)
		for _, professional := range professionals {
			if slices.Contains(professional.ServicesIDS, id.ID(serviceID)) {
				tmp = append(tmp, professional)
			}
		}
		professionals = tmp
	}

	opts := components.WeeklyAppointmentsOptions{
		StartHour:     6,
		EndHour:       20,
		Days:          5,
		Services:      testdata.Services,
		Professionals: professionals,
		Appointments:  appointments,
	}

	if serviceID != "all" {
		professionalID = "all"
	}

	components.WeeklyAppointments("2024-10-10", serviceID, professionalID, opts).Render(r.Context(), w)
}
