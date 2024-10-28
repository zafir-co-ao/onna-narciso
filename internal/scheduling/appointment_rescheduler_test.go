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
		i += 1
		v := strconv.Itoa(i)
		a := scheduling.Appointment{ID: scheduling.NewID(v), Status: scheduling.StatusScheduled}
		repo.Save(a)
	}

	a2 := scheduling.Appointment{ID: "20", Status: scheduling.StatusCanceled}
	a3 := scheduling.Appointment{ID: "12", Status: scheduling.StatusScheduled, Date: "2024-10-27", Start: "8:00", Duration: 240}

	repo.Save(a2)
	repo.Save(a3)

	t.Run("should_reschedule_appointment", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:        "1",
			Date:      "2024-04-11",
			StartHour: "8:30",
			Duration:  120,
		}
		usecase := scheduling.NewAppointmentRescheduler(repo)

		o, err := usecase.Execute(i)
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
		i := scheduling.AppointmentReschedulerInput{
			ID:        "3",
			Date:      "2020-10-10",
			StartHour: "9:30",
			Duration:  120,
		}
		usecase := scheduling.NewAppointmentRescheduler(repo)

		o, err := usecase.Execute(i)
		if err != nil {
			t.Errorf("Should not return error, got %v", err)
		}

		if o.Date != "2020-10-10" {
			t.Errorf("The appointment date must be 2020-10-10, got %v", o.Date)
		}
	})

	t.Run("must_enter_the_reschedule_hour", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:        "4",
			Date:      "2021-11-10",
			StartHour: "10:00",
			Duration:  120,
		}
		usecase := scheduling.NewAppointmentRescheduler(repo)

		o, err := usecase.Execute(i)
		if err != nil {
			t.Errorf("Should not return error, got %v", err)
		}

		if o.Hour != "10:00" {
			t.Errorf("The appointment hour must be 10:00, got %v", o.Hour)
		}
	})

	t.Run("must_enter_the_reschedule_duration", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:        "5",
			Date:      "2022-12-10",
			StartHour: "8:00",
			Duration:  60,
		}

		usecase := scheduling.NewAppointmentRescheduler(repo)

		o, err := usecase.Execute(i)
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

	t.Run("should_return_an_error_if_there_is_no_availability_to_reschedule_the_appointment", func(t *testing.T) {
		var inputs = []scheduling.AppointmentReschedulerInput{
			{
				ID:        "8",
				Date:      "2024-10-27",
				StartHour: "8:00",
				Duration:  60,
			},
			{
				ID:        "9",
				Date:      "2024-10-27",
				StartHour: "9:30",
				Duration:  60,
			},
			{
				ID:        "10",
				Date:      "2024-10-27",
				StartHour: "11:00",
				Duration:  60,
			},
			{
				ID:        "10",
				Date:      "2024-10-27",
				StartHour: "7:00",
				Duration:  90,
			},
		}

		usecase := scheduling.NewAppointmentRescheduler(repo)

		for _, i := range inputs {
			_, err := usecase.Execute(i)

			if errors.Is(nil, err) {
				t.Errorf("Shoud return an error, got %v", err)
			}

			if !errors.Is(scheduling.ErrBusyTime, err) {
				t.Errorf("The error must be ErrBusyTime, got %v", err)
			}
		}
	})

	t.Run("should_return_error_when_appointment_not_found_in_repository", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:        "1000",
			Date:      "2023-11-31",
			StartHour: "9:15",
			Duration:  120,
		}

		usecase := scheduling.NewAppointmentRescheduler(repo)

		_, err := usecase.Execute(i)
		if err == nil {
			t.Errorf("Should return an error, got %v", err)
		}

		if !errors.Is(scheduling.ErrAppointmentNotFound, err) {
			t.Errorf("The error must be ErrAppointmentNotFound, got %v", err)
		}
	})

	t.Run("should_return_error_if_appointment_status_is_different_of_scheduled", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:        "20",
			Date:      "2010-09-18",
			StartHour: "19:15",
			Duration:  60,
		}

		usecase := scheduling.NewAppointmentRescheduler(repo)

		_, err := usecase.Execute(i)
		if err == nil {
			t.Errorf("Should return an error, got %v", err)
		}

		if !errors.Is(scheduling.ErrInvalidStatusToReschedule, err) {
			t.Errorf("The error must be ErrInvalidStatusToReschedule, got %v", err)
		}
	})

	t.Run("should_return_an_error_if_the_date_is_in_an_invalid_format", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:        "6",
			Date:      "10-07-2001",
			StartHour: "11:00",
			Duration:  30,
		}
		usecase := scheduling.NewAppointmentRescheduler(repo)

		_, err := usecase.Execute(i)
		if errors.Is(nil, err) {
			t.Errorf("Shoud return an error, got %v", err)
		}

		if !errors.Is(scheduling.ErrInvalidDate, err) {
			t.Errorf("The error must be ErrInvalidDate, got %v", err)
		}
	})

	t.Run("should_return_an_error_if_the_hour_is_in_an_invalid_format", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:        "2",
			Date:      "2021-07-01",
			StartHour: "11h00",
			Duration:  30,
		}
		usecase := scheduling.NewAppointmentRescheduler(repo)

		_, err := usecase.Execute(i)
		if errors.Is(nil, err) {
			t.Errorf("Shoud return an error, got %v", err)
		}

		if !errors.Is(scheduling.ErrInvalidHour, err) {
			t.Errorf("The error must be ErrHourDate, got %v", err)
		}
	})
}