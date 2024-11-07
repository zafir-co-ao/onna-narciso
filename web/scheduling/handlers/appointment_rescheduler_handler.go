package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	testdata "github.com/zafir-co-ao/onna-narciso/test_data"

	"github.com/zafir-co-ao/onna-narciso/web/scheduling/pages"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleRescheduleAppointment(
	re scheduling.AppointmentRescheduler,
	wg scheduling.WeeklyAppointmentsFinder,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			_http.SendMethodNotAllowed(w)
			return
		}

		duration, err := strconv.Atoi(r.FormValue("duration"))
		if err != nil {
			_http.SendBadRequest(w, "A duração da marcação está no formato inválido")
			return
		}

		input := scheduling.AppointmentReschedulerInput{
			ID:        r.Form.Get("id"),
			Date:      r.Form.Get("date"),
			StartHour: r.Form.Get("start"),
			Duration:  duration,
		}

		_, err = re.Reschedule(input)
		if errors.Is(scheduling.ErrInvalidStatusToReschedule, err) {
			_http.SendBadRequest(w, "Estado inválido para reagendar")
			return
		}

		if errors.Is(scheduling.ErrBusyTime, err) {
			_http.SendBadRequest(w, "Horário Indisponível")
			return
		}

		if errors.Is(scheduling.ErrInvalidDate, err) {
			_http.SendBadRequest(w, "A data para a marcação está no formato inválido")
			return
		}

		if errors.Is(scheduling.ErrInvalidHour, err) {
			_http.SendBadRequest(w, "A hora da marcação está no formato inválido")
			return
		}

		if errors.Is(scheduling.ErrAppointmentNotFound, err) {
			_http.SendNotFound(w, "Marcação não encontrada")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		weekDay := r.FormValue("week-day")
		serviceID := r.FormValue("service-id")
		professionalID := r.FormValue("professional-id")

		appointments, err := wg.Find(
			weekDay,
			serviceID,
			[]string{professionalID},
		)

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		//TODO - Utilizar o repositório de profissionais para filtrar os profissionais que atendem o serviço
		professionals := testdata.FindProfessionalsByServiceID(serviceID)

		_http.SendOk(w)

		opts := pages.WeeklyAppointmentsOptions{
			ServiceID:      serviceID,
			ProfessionalID: professionalID,
			Services:       testdata.Services,
			Date:           weekDay,
			Days:           5,
			StartHour:      6,
			EndHour:        20,
			Appointments:   appointments,
			Professionals:  professionals,
		}

		pages.WeeklyAppointments(opts).Render(r.Context(), w)
	}
}
