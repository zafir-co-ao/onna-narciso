package scheduling_test

import (
	"errors"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"
)

func TestAppointmentScheduler(t *testing.T) {
	t.Run("should_schedule_appointment", func(t *testing.T) {
		repo := inmem.NewAppointmentRepository()
		usecase := scheduling.NewAppointmentScheduler(repo)

		id, err := usecase.Schedule()

		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if id != "1" {
			t.Error("Appointment id should be 1")
		}
	})

	t.Run("should_store_appointment_in_repository", func(t *testing.T) {

		repo := inmem.NewAppointmentRepository()

		usecase := scheduling.NewAppointmentScheduler(repo)

		id, err := usecase.Schedule()
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.Get(id)
		if errors.Is(err, scheduling.ErrAppointmentNotFound) {
			t.Error("Appointment should be stored in repository")
		}

		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.ID != id {
			t.Errorf("Appointment ID should be %s", id)
		}
	})

	t.Run("the_status_of_appointment_should_be_scheduled", func(t *testing.T) {
		repo := inmem.NewAppointmentRepository()

		usecase := scheduling.NewAppointmentScheduler(repo)

		id, err := usecase.Schedule()
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.Get(id)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.Status != "Scheduled" {
			t.Errorf("Appointment status should be Scheduled, got %s", appointment.Status)
		}
	})
}
