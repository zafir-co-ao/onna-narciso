package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/web/components"
)

func HandleAppointmentForm(w http.ResponseWriter, r *http.Request) {
	var s = components.FormState{
		ProfessionalID:   "1",
		ProfessionalName: "Sara Gomes",
		ServiceID:        "1",
		ServiceName:      "Manicure + Pedicure",
		Hour:             "10:00",
		Date:             "2024-10-10",
	}
	components.ScheduleEventForm(s).Render(r.Context(), w)
}
