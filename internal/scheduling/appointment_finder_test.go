package scheduling_test

import (
	"errors"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"
)

func TestAppointmentFinder(t *testing.T) {
	repo := inmem.NewAppointmentRepository()

	repo.Save(scheduling.Appointment{ID: "1"})

	t.Run("should_find_appointment_in_repository", func(t *testing.T) {
		usecase := scheduling.NewAppointmentFinder(repo)

		a, err := usecase.Execute("1")
		if err != nil {
			t.Errorf("Finder appointment should not return error: %v", err)
		}

		if errors.Is(scheduling.ErrAppointmentNotFound, err) {
			t.Errorf("Should return appointment not an error, got %v", err)
		}

		if a.ID != "1" {
			t.Errorf("Should return appointment with id %v, got %v", 1, a.ID)
		}
	})

	t.Run("should_return_error_when_appointment_not_found_in_repository", func(t *testing.T) {
		usecase := scheduling.NewAppointmentFinder(repo)

		_, err := usecase.Execute("2")
		if err == nil {
			t.Errorf("Finder appointment should return error, got %v", err)
		}

		if !errors.Is(scheduling.ErrAppointmentNotFound, err) {
			t.Errorf("The error must be ErrAppointmentNotFound, got %v", err)
		}
	})
}
