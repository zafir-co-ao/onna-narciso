package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/web/components"
)

func HandleAppointmentForm(w http.ResponseWriter, r *http.Request) {
	var s = components.AppointmentSchedulingState{
		ProfessionalID:   r.FormValue("professional_id"),
		ProfessionalName: r.FormValue("professional_name"),
		ServiceID:        r.FormValue("service_id"),
		ServiceName:      r.FormValue("service_name"),
		StartHour:        r.FormValue("hour"),
		Date:             r.FormValue("date"),
	}
	components.AppointmentSchedulingForm(s).Render(r.Context(), w)
}
