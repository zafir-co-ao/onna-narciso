package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/hour"

	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleRescheduleAppointment(re scheduling.AppointmentRescheduler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		duration, err := strconv.Atoi(r.FormValue("duration"))
		if err != nil {
			_http.SendBadRequest(w, "A duração da marcação está no formato inválido")
			return
		}

		input := scheduling.AppointmentReschedulerInput{
			ID:             r.Form.Get("id"),
			Date:           r.Form.Get("date"),
			Hour:           r.Form.Get("hour"),
			ProfessionalID: r.Form.Get("professional-id"),
			ServiceID:      r.Form.Get("service-id"),
			Duration:       duration,
		}

		_, err = re.Reschedule(input)
		if errors.Is(err, scheduling.ErrInvalidStatusToReschedule) {
			_http.SendBadRequest(w, "Estado inválido para reagendar")
			return
		}

		if errors.Is(err, scheduling.ErrBusyTime) {
			_http.SendBadRequest(w, "Horário Indisponível")
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

		if errors.Is(err, scheduling.ErrInvalidService) {
			_http.SendBadRequest(w, "Indisponibilidade do serviço para o profissional")
			return
		}

		if errors.Is(err, scheduling.ErrAppointmentNotFound) {
			_http.SendNotFound(w, "Marcação não encontrada")
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

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		// weekDay := r.FormValue("week-day")
		// serviceID := r.FormValue("service-id")
		// professionalID := r.FormValue("professional-id")

		// appointments, err := wg.Find(
		// 	weekDay,
		// 	serviceID,
		// 	[]string{professionalID},
		// )

		if err != nil {
			_http.SendServerError(w)
			return
		}

		_http.SendOk(w)

		//TODO - Utilizar o repositório de profissionais para filtrar os profissionais que atendem o serviço
		// professionals := testdata.FindProfessionalsByServiceID(serviceID)

		// opts := pages.WeeklyAppointmentsOptions{
		// 	ServiceID:      serviceID,
		// 	ProfessionalID: professionalID,
		// 	Services:       testdata.Services,
		// 	Date:           weekDay,
		// 	Days:           5,
		// 	StartHour:      6,
		// 	EndHour:        20,
		// 	Appointments:   appointments,
		// 	Professionals:  professionals,
		// }

		// pages.WeeklyAppointments(opts).Render(r.Context(), w)
	}
}
