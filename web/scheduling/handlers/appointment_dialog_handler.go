package handlers

import (
	"net/http"
	"strconv"

	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/components"
)

func HandleScheduleAppointmentDialog() func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		var s = components.AppointmentSchedulerState{
			ProfessionalID: r.FormValue("professional-id"),
			ServiceID:      r.FormValue("service-id"),
			Hour:           r.FormValue("hour"),
			Date:           r.FormValue("date"),
			HxTarget:       r.FormValue("hx-target"),
			HxSwap:         r.FormValue("hx-swap"),
		}

		for _, p := range testdata.Professionals {
			if p.ID.String() == s.ProfessionalID {
				s.ProfessionalName = p.Name.String()
				break
			}
		}

		for _, svc := range testdata.Services {
			if svc.ID.String() == s.ServiceID {
				s.ServiceName = svc.Name.String()
				s.ServiceDuration = strconv.Itoa(svc.Duration)
				break
			}
		}

		components.AppointmentSchedulerDialog(s).Render(r.Context(), w)
	}
}
