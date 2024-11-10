package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/hour"
	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/pages"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleScheduleAppointment(
	s scheduling.AppointmentScheduler,
	wg scheduling.WeeklyAppointmentsFinder,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		duration, err := strconv.Atoi(r.FormValue("duration"))
		if err != nil {
			_http.SendBadRequest(w, "A duração da marcação está no formato inválido")
			return
		}

		input := scheduling.AppointmentSchedulerInput{
			ProfessionalID: r.Form.Get("professional-id"),
			ServiceID:      r.Form.Get("service-id"),
			Date:           r.Form.Get("date"),
			Hour:           r.Form.Get("hour"),
			CustomerID:     r.Form.Get("customer-id"),
			CustomerName:   r.Form.Get("customer-name"),
			CustomerPhone:  r.Form.Get("customer-phone"),
			Duration:       duration,
		}

		_, err = s.Schedule(input)

		if errors.Is(err, scheduling.ErrCustomerNotFound) {
			_http.SendNotFound(w, "Cliente não encontrado")
			return
		}

		if errors.Is(err, scheduling.ErrCustomerRegistration) {
			_http.SendReponse(w, "Não foi possível registar o cliente", http.StatusInternalServerError)
			return
		}

		if errors.Is(err, scheduling.ErrProfessionalNotFound) {
			_http.SendNotFound(w, "Profissional não encontrado")
			return
		}

		if errors.Is(err, scheduling.ErrServiceNotFound) {
			_http.SendNotFound(w, "Serviço não encontrado")
			return
		}

		if errors.Is(err, scheduling.ErrBusyTime) {
			_http.SendBadRequest(w, "Horarário Indisponível")
			return
		}

		if errors.Is(err, date.ErrInvalidDate) {
			_http.SendBadRequest(w, "A data para a marcação está no formato inválido")
			return
		}

		if errors.Is(err, hour.ErrInvalidHour) {
			_http.SendBadRequest(w, "A hora da marcação está no formato inválido")
			return
		}

		if err != nil {
			_http.SendServerError(w)
			return
		}

		appointments, err := wg.Find(
			input.Date,
			input.ServiceID,
			[]string{input.ProfessionalID},
		)

		if err != nil {
			_http.SendServerError(w)
			return
		}

		_http.SendCreated(w)

		//TODO - Utilizar o repositório de profissionais para filtrar os profissionais que atendem o serviço
		professionals := testdata.FindProfessionalsByServiceID(input.ServiceID)

		opts := pages.WeeklyAppointmentsOptions{
			ServiceID:      input.ServiceID,
			ProfessionalID: input.ProfessionalID,
			Date:           input.Date,
			StartHour:      6,
			EndHour:        20,
			Days:           5,
			Appointments:   appointments,
			Services:       testdata.Services,
			Professionals:  professionals,
		}

		pages.WeeklyAppointments(opts).Render(r.Context(), w)
	}
}
