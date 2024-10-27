package scheduling_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"
)

func TestAppointmentRescheduler(t *testing.T) {
	repo := inmem.NewAppointmentRepository()

	for i := range 10 {
		v := strconv.Itoa(i)
		a := scheduling.Appointment{ID: scheduling.NewID(v), Status: scheduling.StatusScheduled}
		repo.Save(a)
	}

	a2 := scheduling.Appointment{ID: "20", Status: scheduling.StatusCanceled}

	repo.Save(a2)

	t.Run("should_reschedule_appointment", func(t *testing.T) {
		usecase := scheduling.NewAppointmentRescheduler(repo)

		o, err := usecase.Execute("1", "2024-04-11", "8:30", 120)
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

		o, err := usecase.Execute("3", "2020-10-10", "9:30", 120)
		if err != nil {
			t.Errorf("Should not return error, got %v", err)
		}

		if o.Date != "2020-10-10" {
			t.Errorf("The appointment date must be 2020-10-10, got %v", o.Date)
		}
	})

	t.Run("must_enter_the_reschedule_hour", func(t *testing.T) {
		usecase := scheduling.NewAppointmentRescheduler(repo)

		o, err := usecase.Execute("4", "2021-11-10", "10:00", 120)
		if err != nil {
			t.Errorf("Should not return error, got %v", err)
		}

		if o.Hour != "10:00" {
			t.Errorf("The appointment hour must be 10:00, got %v", o.Hour)
		}
	})

	t.Run("must_enter_the_reschedule_duration", func(t *testing.T) {
		usecase := scheduling.NewAppointmentRescheduler(repo)

		o, err := usecase.Execute("5", "2022-12-10", "8:00", 60)
		if err != nil {
			t.Errorf("Should not return error, got %v", err)
		}

		a, err := repo.FindByID(scheduling.NewID(o.ID))
		if err != nil {
			t.Errorf("Should return the appointment not an error, got %v", err)
		}

		if o.Duration != 60 {
			t.Errorf("The appointment duration must be 60 minutes, got %v", o.Duration)
		}

		if a.End.Value() != "9:00" {
			t.Errorf("The appointment end must be 9:00, got %v", o.Duration)
		}
	})

	t.Run("should_return_error_when_appointment_not_found_in_repository", func(t *testing.T) {
		usecase := scheduling.NewAppointmentRescheduler(repo)

		_, err := usecase.Execute("1000", "2023-11-31", "9:15", 120)
		if err == nil {
			t.Errorf("Should return an error, got %v", err)
		}

		if !errors.Is(scheduling.ErrAppointmentNotFound, err) {
			t.Errorf("The error must be ErrAppointmentNotFound, got %v", err)
		}
	})

	t.Run("should_return_error_if_appointment_status_is_different_of_scheduled", func(t *testing.T) {
		usecase := scheduling.NewAppointmentRescheduler(repo)

		_, err := usecase.Execute("20", "2010-09-18", "19:15", 120)
		if err == nil {
			t.Errorf("Should return an error, got %v", err)
		}

		if !errors.Is(scheduling.ErrInvalidStatusToReschedule, err) {
			t.Errorf("The error must be ErrInvalidStatusToReschedule, got %v", err)
		}
	})

	t.Run("should_return_an_error_if_the_date_is_in_an_invalid_format", func(t *testing.T) {
		usecase := scheduling.NewAppointmentRescheduler(repo)

		_, err := usecase.Execute("6", "10-10-2020", "10:00", 120)
		if errors.Is(nil, err) {
			t.Errorf("Shoud return an error, got %v", err)
		}

		if !errors.Is(scheduling.ErrInvalidDate, err) {
			t.Errorf("The error must be ErrInvalidDate, got %v", err)
		}
	})

	t.Run("should_return_an_error_if_the_hour_is_in_an_invalid_format", func(t *testing.T) {
		usecase := scheduling.NewAppointmentRescheduler(repo)

		_, err := usecase.Execute("2", "2021-07-01", "11h00", 30)
		if errors.Is(nil, err) {
			t.Errorf("Shoud return an error, got %v", err)
		}

		if !errors.Is(scheduling.ErrInvalidHour, err) {
			t.Errorf("The error must be ErrHourDate, got %v", err)
		}
	})
}
