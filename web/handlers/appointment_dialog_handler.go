package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/web/components"
)

func HandleAppointmentDialog(w http.ResponseWriter, r *http.Request) {
	var s = components.AppointmentSchedulerState{
		ProfessionalID:   r.FormValue("professional-id"),
		ProfessionalName: r.FormValue("professional-name"),
		ServiceID:        r.FormValue("service-id"),
		ServiceName:      r.FormValue("service-name"),
		StartHour:        r.FormValue("hour"),
		Date:             r.FormValue("date"),
		HxTarget:         r.FormValue("hx-target"),
		HxSwap:           r.FormValue("hx-swap"),
	}
	components.AppointmentSchedulerDialog(s).Render(r.Context(), w)
}
