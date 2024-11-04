package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/zafir-co-ao/onna-narciso/web/scheduling/components"
)

func HandleAppointmentDialog(w http.ResponseWriter, r *http.Request) {
	startHour := 6
	endHourt := 20
	hours := (endHourt-startHour)*4 + 1
	height, _ := strconv.Atoi(r.FormValue("height"))
	layerY, _ := strconv.Atoi(r.FormValue("layer-y"))

	hour := getStartHour(height, layerY, hours, startHour)

	var s = components.AppointmentSchedulerState{
		ProfessionalID:   r.FormValue("professional-id"),
		ProfessionalName: r.FormValue("professional-name"),
		ServiceID:        r.FormValue("service-id"),
		ServiceName:      r.FormValue("service-name"),
		StartHour:        hour,
		Date:             r.FormValue("date"),
		HxTarget:         r.FormValue("hx-target"),
		HxSwap:           r.FormValue("hx-swap"),
		HxTrigger:        r.FormValue("hx-trigger"),
	}

	components.AppointmentSchedulerDialog(s).Render(r.Context(), w)
}

func getStartHour(height, layerY, hours, startHour int) string {
	h := calculateHour(height, layerY, hours, startHour)
	m := calculateMinutes(height, layerY, hours, startHour)
	if len(m) == 1 {
		return fmt.Sprintf("%s:0%s", h, m)
	}

	return fmt.Sprintf("%s:%s", h, m)
}

func calculateHour(height, layerY, hours, startHour int) string {
	row := height / hours
	d := (layerY - (3 * row)) / row
	h := d/4 + startHour
	return fmt.Sprintf("%d", h)
}

func calculateMinutes(height, layerY, hours, startHour int) string {
	row := height / hours
	d := (layerY - (3 * row)) / row
	m := d % 4 * 15
	return fmt.Sprintf("%d", m)
}
