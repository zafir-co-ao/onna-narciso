package handlers

import (
	"errors"
	"net/http"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
	"github.com/zafir-co-ao/onna-narciso/web/scheduling/pages"
	_http "github.com/zafir-co-ao/onna-narciso/web/shared/http"
)

func HandleCancelAppointment(
	c scheduling.AppointmentCanceler,
	wg scheduling.WeeklyAppointmentsFinder,
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := c.Execute(r.PathValue("id"))

		if errors.Is(scheduling.ErrInvalidStatusToCancel, err) {
			_http.SendBadRequest(w, "Estado inválido para cancelar")
			return
		}

		if errors.Is(scheduling.ErrAppointmentNotFound, err) {
			_http.SendBadRequest(w, "Marcação não encontrada")
			return
		}

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		weekDay := r.FormValue("week-day")
		serviceID := r.FormValue("service-id")
		professionalID := r.FormValue("professional-id")

		appointments, err := wg.Find(weekDay, serviceID, []string{professionalID})

		if !errors.Is(nil, err) {
			_http.SendServerError(w)
			return
		}

		_http.SendOk(w)

		//TODO - Utilizar o repositório de profissionais para filtrar os profissionais que atendem o serviço
		professionals := testdata.FindProfessionalsByServiceID(serviceID)

		opts := pages.WeeklyAppointmentsOptions{
			Date:           weekDay,
			ServiceID:      serviceID,
			ProfessionalID: professionalID,
			EndHour:        20,
			StartHour:      6,
			Days:           5,
			Services:       testdata.Services,
			Professionals:  professionals,
			Appointments:   appointments,
		}

		pages.WeeklyAppointments(opts).Render(r.Context(), w)
	}
}
