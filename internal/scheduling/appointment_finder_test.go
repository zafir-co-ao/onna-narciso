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
	repo.Save(scheduling.Appointment{ID: "2"})

	t.Run("should_find_appointment_in_repository", func(t *testing.T) {
		usecase := scheduling.NewAppointmentFinder(repo)

		a, err := usecase.Execute("1")
		if err != nil {
			t.Errorf("Finder appointment should not return error: %v", err)
		}

		if errors.Is(scheduling.ErrAppointmentNotFound, err) {
			t.Errorf("Should return appointment not an error: %v", err)
		}

		if a.ID.Value() != "1" {
			t.Errorf("Should return appointment with id %v but got %v", 1, a.ID.Value())
		}
	})
}
