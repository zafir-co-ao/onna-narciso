package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/web/components"
)

var appointments = []scheduling.Appointment{
	{
		ID:               "1",
		CustomerName:     "Paola Oliveira",
		ProfessionalName: "Julieta Venegas",
		ServiceName:      "Manicure",
		Date:             scheduling.Date("2024-10-10"),
		Start:            scheduling.Hour("08:00"),
		Duration:         180,
	},
	{
		ID:               "2",
		CustomerName:     "Juliana Paes",
		ProfessionalName: "Julieta Venegas",
		ServiceName:      "Manicure",
		Date:             scheduling.Date("2024-10-11"),
		Start:            scheduling.Hour("10:30"),
		Duration:         90,
	},
	{
		ID:               "3",
		CustomerName:     "Gisele Bündchen",
		ProfessionalName: "Mariana Aydar",
		ServiceName:      "Depilação Laser",
		Date:             scheduling.Date("2024-10-10"),
		Start:            scheduling.Hour("12:00"),
		Duration:         60,
	},
}

func HandleWeeklyAppointments(w http.ResponseWriter, r *http.Request) {
	components.WeeklyAppointments("2024-10-10", 6, 8, 22, appointments).Render(r.Context(), w)
}
