package scheduling_test

import (
	"errors"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"
)

func TestAppointmentCanceler(t *testing.T) {
	a1, _ := scheduling.NewAppointment("1", "10", "Jane Doe", "3", "Sara Gomes", "4", "2021-01-01", "10:00", 60)
	a2, _ := scheduling.NewAppointment("1", "10", "Jane Doe", "4", "Micheal Miller", "4", "2021-01-01", "10:00", 60)

	t.Run("should_cancel_an_appointment", func(t *testing.T) {
		repo := inmem.NewAppointmentRepository()
		repo.Save(a1)

		usecase := scheduling.NewAppointmentCanceler(repo)

		err := usecase.Execute("1")
		if err != nil {
			t.Errorf("Canceling appointment should not return error: %v", err)
		}

		app, err := repo.FindByID("1")
		if errors.Is(scheduling.ErrAppointmentNotFound, err) {
			t.Errorf("Appointment should be stored in repository")
		}

		if !app.IsCancelled() {
			t.Errorf("Appointment should be status cancelled, got: %v", app.Status)
		}
	})

	t.Run("should_return_error_if_appointment_status_is_canceled", func(t *testing.T) {
		repo := inmem.NewAppointmentRepository()
		repo.Save(a2)

		usecase := scheduling.NewAppointmentCanceler(repo)

		err := usecase.Execute("1")
		if err != nil {
			t.Errorf("Canceling appointment should not return error: %v", err)
		}

		err = usecase.Execute("1")
		if err == nil {
			t.Errorf("Canceling appointment should return error: %v", err)
		}

		if !errors.Is(scheduling.ErrInvalidStatusToCancel, err) {
			t.Errorf("Canceling appointment should return error: %v", err)
		}
	})

	t.Run("should_return_error_appointment_not_found_when_appointment_not_exists_in_repository", func(t *testing.T) {
		repo := inmem.NewAppointmentRepository()
		usecase := scheduling.NewAppointmentCanceler(repo)

		err := usecase.Execute("1")
		if err == nil {
			t.Errorf("Canceling appointment should return error: %v", err)
		}

		if !errors.Is(scheduling.ErrAppointmentNotFound, err) {
			t.Errorf("Canceling appointment should return error: %v", err)
		}
	})
}
