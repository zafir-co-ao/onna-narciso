package scheduling_test

import (
	"errors"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
)

func TestAppointmentGetter(t *testing.T) {
	a := scheduling.Appointment{ID: "1"}
	repo := scheduling.NewAppointmentRepository(a)
	u := scheduling.NewAppointmentGetter(repo)

	t.Run("should_retrieve_appointment_from_repository", func(t *testing.T) {
		a, err := u.Get("1")

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if errors.Is(err, scheduling.ErrAppointmentNotFound) {
			t.Errorf("Should return the appointment, got %v", err)
		}

		if a.ID != "1" {
			t.Errorf("Should return appointment with id %v, got %v", 1, a.ID)
		}
	})

	t.Run("should_return_error_when_appointment_not_found_in_repository", func(t *testing.T) {
		_, err := u.Get("2")

		if errors.Is(nil, err) {
			t.Errorf("Expected an error, got %v", err)
		}

		if !errors.Is(err, scheduling.ErrAppointmentNotFound) {
			t.Errorf("The error must be %v, got %v", scheduling.ErrAppointmentNotFound, err)
		}
	})
}
