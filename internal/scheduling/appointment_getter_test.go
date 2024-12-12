package scheduling_test

import (
	"errors"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
)

func TestAppointmentGetter(t *testing.T) {
	repo := scheduling.NewAppointmentRepository()
	repo.Save(scheduling.Appointment{ID: "1"})

	u := scheduling.NewAppointmentGetter(repo)

	t.Run("should_find_appointment_in_repository", func(t *testing.T) {
		a, err := u.Get("1")
		if err != nil {
			t.Errorf("Finder appointment should not return error: %v", err)
		}

		if errors.Is(err, scheduling.ErrAppointmentNotFound) {
			t.Errorf("Should return appointment not an error, got %v", err)
		}

		if a.ID != "1" {
			t.Errorf("Should return appointment with id %v, got %v", 1, a.ID)
		}
	})

	t.Run("should_return_error_when_appointment_not_found_in_repository", func(t *testing.T) {
		_, err := u.Get("2")
		if err == nil {
			t.Errorf("Finder appointment should return error, got %v", err)
		}

		if !errors.Is(err, scheduling.ErrAppointmentNotFound) {
			t.Errorf("The error must be ErrAppointmentNotFound, got %v", err)
		}
	})
}
