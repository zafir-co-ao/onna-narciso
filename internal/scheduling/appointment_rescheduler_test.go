package scheduling_test

import (
	"errors"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"
)

func TestAppointmentRescheduler(t *testing.T) {
	repo := inmem.NewAppointmentRepository()

	a := scheduling.Appointment{ID: "1"}
	repo.Save(a)

	t.Run("should_reschedule_appointment", func(t *testing.T) {
		usecase := scheduling.NewAppointmentRescheduler(repo)

		o, err := usecase.Execute("1")
		if err != nil {
			t.Errorf("Should not return error, got %v", err)
		}

		a, err := repo.FindByID(scheduling.NewID(o.ID))
		if err != nil {
			t.Errorf("Should return the appointment not an error, got %v", err)
		}

		if !a.IsRescheduled() {
			t.Errorf("The appointment status must be Rescheduled, got %v", a.Status)
		}
	})

	t.Run("should_return_error_when_appointment_not_found_in_repository", func(t *testing.T) {
		usecase := scheduling.NewAppointmentRescheduler(repo)

		_, err := usecase.Execute("2")
		if err == nil {
			t.Errorf("Should return an error, got %v", err)
		}

		if !errors.Is(scheduling.ErrAppointmentNotFound, err) {
			t.Errorf("The error must be ErrAppointmentNotFound, got %v", err)
		}
	})
}
