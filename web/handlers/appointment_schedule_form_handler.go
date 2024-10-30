package handlers

import (
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/web/components"
)

func HandleAppointmentScheduleForm(w http.ResponseWriter, r *http.Request) {
	components.ScheduleEventForm().Render(r.Context(), w)
}
