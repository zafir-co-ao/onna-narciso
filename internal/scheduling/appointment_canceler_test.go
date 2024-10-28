package scheduling_test

import (
	"errors"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"
)

func TestAppointmentCanceler(t *testing.T) {
	repo := inmem.NewAppointmentRepository()
	a1 := scheduling.Appointment{ID: "1"}
	a2 := scheduling.Appointment{ID: "2", Status: scheduling.StatusCanceled}

	repo.Save(a1)
	repo.Save(a2)

	t.Run("should_cancel_an_appointment", func(t *testing.T) {
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
		usecase := scheduling.NewAppointmentCanceler(repo)

		err := usecase.Execute("2")
		if err == nil {
			t.Errorf("Canceling appointment should return error: %v", err)
		}

		if !errors.Is(scheduling.ErrInvalidStatusToCancel, err) {
			t.Errorf("Canceling appointment should return error: %v", err)
		}
	})

	t.Run("should_return_error_appointment_not_found_when_appointment_not_exists_in_repository", func(t *testing.T) {
		usecase := scheduling.NewAppointmentCanceler(repo)

		err := usecase.Execute("3")
		if err == nil {
			t.Errorf("Canceling appointment should return error: %v", err)
		}

		if !errors.Is(scheduling.ErrAppointmentNotFound, err) {
			t.Errorf("Canceling appointment should return error: %v", err)
		}
	})
}
