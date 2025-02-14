package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/hr"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"

	"github.com/zafir-co-ao/onna-narciso/web/scheduling/components"
)

func HandleScheduleAppointmentDialog(cf crm.CustomerFinder, pf hr.ProfessionalFinder, sf services.ServiceFinder) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		customers, err := cf.FindAll()
		if !errors.Is(err, nil) {
			_http.SendServerError(w)
			return
		}

		p, err := pf.FindByID(r.FormValue("professional-id"))

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		if errors.Is(err, hr.ErrProfessionalNotFound) {
			_http.SendNotFound(w, "Profissional não encontrado")
			return
		}

		s, err := sf.FindByID(r.FormValue("service-id"))

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		if errors.Is(err, services.ErrServiceNotFound) {
			_http.SendNotFound(w, "Serviço não encontrado")
			return
		}

		var as = components.AppointmentSchedulerState{
			ProfessionalID:   p.ID,
			ProfessionalName: p.Name,
			ServiceID:        s.ID,
			ServiceName:      s.Name,
			ServiceDuration:  strconv.Itoa(s.Duration),
			Hour:             r.FormValue("hour"),
			Date:             r.FormValue("date"),
			HxPost:           r.FormValue("hx-post"),
			Customers:        customers,
		}

		_http.SendOk(w)
		components.AppointmentSchedulerDialog(as).Render(r.Context(), w)
	}
}
