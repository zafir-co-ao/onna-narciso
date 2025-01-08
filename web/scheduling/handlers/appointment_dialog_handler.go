package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"

	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/components"
)

func HandleScheduleAppointmentDialog(cf crm.CustomerFinder) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		customers, err := cf.FindAll()
		if !errors.Is(err, nil) {
			_http.SendServerError(w)
			return
		}

		var s = components.AppointmentSchedulerState{
			ProfessionalID: r.FormValue("professional-id"),
			ServiceID:      r.FormValue("service-id"),
			Hour:           r.FormValue("hour"),
			Date:           r.FormValue("date"),
			HxPost:         r.FormValue("hx-post"),
			Customers:      customers,
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
				s.ServiceDuration = svc.Duration.String()
				break
			}
		}

		components.AppointmentSchedulerDialog(s).Render(r.Context(), w)
	}
}
