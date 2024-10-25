package scheduling_test

import (
	"errors"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"
)

func TestAppointmentRescheduler(t *testing.T) {
	repo := inmem.NewAppointmentRepository()

	a1 := scheduling.Appointment{ID: "1", Status: scheduling.StatusScheduled}
	a2 := scheduling.Appointment{ID: "2", Status: scheduling.StatusCanceled}
	a3 := scheduling.Appointment{ID: "3", Status: scheduling.StatusScheduled}
	a4 := scheduling.Appointment{ID: "4", Status: scheduling.StatusScheduled}

	repo.Save(a1)
	repo.Save(a2)
	repo.Save(a3)
	repo.Save(a4)

	t.Run("should_reschedule_appointment", func(t *testing.T) {
		usecase := scheduling.NewAppointmentRescheduler(repo)

		o, err := usecase.Execute("1", "2024-04-11", "8:30")
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

	t.Run("must_enter_the_reschedule_date", func(t *testing.T) {
		usecase := scheduling.NewAppointmentRescheduler(repo)

		o, err := usecase.Execute("3", "2020-10-10", "9:30")
		if err != nil {
			t.Errorf("Should not return error, got %v", err)
		}

		if o.Date != "2020-10-10" {
			t.Errorf("The appointment date must be 2020-10-10, got %v", o.Date)
		}
	})

	t.Run("must_enter_the_reschedule_hour", func(t *testing.T) {
		usecase := scheduling.NewAppointmentRescheduler(repo)

		o, err := usecase.Execute("4", "2021-11-10", "10:00")
		if err != nil {
			t.Errorf("Should not return error, got %v", err)
		}

		if o.Hour != "10:00" {
			t.Errorf("The appointment hour must be 10:00, got %v", o.Hour)
		}
	})

	t.Run("should_return_error_when_appointment_not_found_in_repository", func(t *testing.T) {
		usecase := scheduling.NewAppointmentRescheduler(repo)

		_, err := usecase.Execute("1000", "2023-11-31", "9:15")
		if err == nil {
			t.Errorf("Should return an error, got %v", err)
		}

		if !errors.Is(scheduling.ErrAppointmentNotFound, err) {
			t.Errorf("The error must be ErrAppointmentNotFound, got %v", err)
		}
	})

	t.Run("should_return_error_if_appointment_status_is_different_of_scheduled", func(t *testing.T) {
		usecase := scheduling.NewAppointmentRescheduler(repo)

		_, err := usecase.Execute("2", "2010-09-18", "19:15")
		if err == nil {
			t.Errorf("Should return an error, got %v", err)
		}

		if !errors.Is(scheduling.ErrInvalidStatusToReschedule, err) {
			t.Errorf("The error must be ErrInvalidStatusToReschedule, got %v", err)
		}
	})
}
