package scheduling

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func NewAppointmentSchedulerHandler(u AppointmentScheduler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		duration, err := strconv.ParseInt(r.FormValue("duration"), 10, 64)
		if err != nil {
			http.Error(w, "invalid duration", http.StatusBadRequest)
			return
		}

		input := AppointmentSchedulerInput{
			ProfessionalID: r.Form.Get("professional_id"),
			ServiceID:      r.Form.Get("service_id"),
			Date:           r.Form.Get("date"),
			StartHour:      r.Form.Get("start"),
			CustomerID:     r.Form.Get("customer_id"),
			CustomerName:   r.Form.Get("customer_name"),
			CustomerPhone:  r.Form.Get("customer_phone"),
			Duration:       int(duration),
		}

		id, err := u.Schedule(input)

		if errors.Is(ErrCustomerRegistration, err) {
			http.Error(w, "customer registration error", http.StatusInternalServerError)
			return
		}

		if errors.Is(ErrCustomerNotFound, err) {
			http.Error(w, "customer not found", http.StatusNotFound)
			return
		}

		if errors.Is(ErrProfessionalNotFound, err) {
			http.Error(w, "professional not found", http.StatusNotFound)
			return
		}

		if errors.Is(ErrServiceNotFound, err) {
			http.Error(w, "service not found", http.StatusNotFound)
			return
		}

		if errors.Is(ErrBusyTime, err) {
			http.Error(w, "busy time", http.StatusBadRequest)
			return
		}

		if errors.Is(ErrInvalidDate, err) {
			http.Error(w, "invalid date", http.StatusBadRequest)
			return
		}

		if errors.Is(ErrInvalidHour, err) {
			http.Error(w, "invalid start hour", http.StatusBadRequest)
			return
		}

		if !errors.Is(nil, err) {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf("appointment scheduled id: %v", id)))
	}
}
