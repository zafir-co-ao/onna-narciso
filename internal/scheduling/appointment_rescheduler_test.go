package scheduling_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/event"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/hour"
)

func TestAppointmentRescheduler(t *testing.T) {
	bus := event.NewInmemEventBus()
	repo := inmem.NewAppointmentRepository()

	for i := range 20 {
		i += 1
		v := strconv.Itoa(i)
		a := scheduling.Appointment{ID: nanoid.ID(v), Status: scheduling.StatusScheduled}
		repo.Save(a)
	}

	a2 := scheduling.Appointment{ID: "20", Status: scheduling.StatusCanceled}
	a3 := scheduling.Appointment{ID: "12", Status: scheduling.StatusScheduled, Date: "2024-10-27", Hour: "8:00", Duration: 240}

	repo.Save(a2)
	repo.Save(a3)

	t.Run("should_reschedule_appointment", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:       "1",
			Date:     "2024-04-11",
			Hour:     "8:30",
			Duration: 120,
		}
		usecase := scheduling.NewAppointmentRescheduler(repo, bus)

		o, err := usecase.Reschedule(i)
		if err != nil {
			t.Errorf("Should not return error, got %v", err)
		}

		a, err := repo.FindByID(nanoid.ID(o.ID))
		if err != nil {
			t.Errorf("Should return the appointment not an error, got %v", err)
		}

		if !a.IsRescheduled() {
			t.Errorf("The appointment status must be Rescheduled, got %v", a.Status)
		}
	})

	t.Run("must_enter_the_reschedule_date", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:       "3",
			Date:     "2020-10-10",
			Hour:     "9:30",
			Duration: 120,
		}
		usecase := scheduling.NewAppointmentRescheduler(repo, bus)

		o, err := usecase.Reschedule(i)
		if err != nil {
			t.Errorf("Should not return error, got %v", err)
		}

		if o.Date != "2020-10-10" {
			t.Errorf("The appointment date must be 2020-10-10, got %v", o.Date)
		}
	})

	t.Run("must_enter_the_reschedule_hour", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:       "4",
			Date:     "2021-11-10",
			Hour:     "10:00",
			Duration: 120,
		}
		usecase := scheduling.NewAppointmentRescheduler(repo, bus)

		o, err := usecase.Reschedule(i)
		if err != nil {
			t.Errorf("Should not return error, got %v", err)
		}

		if o.Hour != "10:00" {
			t.Errorf("The appointment hour must be 10:00, got %v", o.Hour)
		}
	})

	t.Run("must_enter_the_reschedule_duration", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:       "5",
			Date:     "2022-12-10",
			Hour:     "8:00",
			Duration: 60,
		}

		usecase := scheduling.NewAppointmentRescheduler(repo, bus)

		o, err := usecase.Reschedule(i)
		if err != nil {
			t.Errorf("Should not return error, got %v", err)
		}

		_, err = repo.FindByID(nanoid.ID(o.ID))
		if err != nil {
			t.Errorf("Should return the appointment not an error, got %v", err)
		}

		if o.Duration != 60 {
			t.Errorf("The appointment duration must be 60 minutes, got %v", o.Duration)
		}

	})

	t.Run("should_return_an_error_if_there_is_no_availability_to_reschedule_the_appointment", func(t *testing.T) {
		var inputs = []scheduling.AppointmentReschedulerInput{
			{
				ID:       "8",
				Date:     "2024-10-27",
				Hour:     "8:00",
				Duration: 60,
			},
			{
				ID:       "9",
				Date:     "2024-10-27",
				Hour:     "9:30",
				Duration: 60,
			},
			{
				ID:       "10",
				Date:     "2024-10-27",
				Hour:     "11:00",
				Duration: 60,
			},
			{
				ID:       "10",
				Date:     "2024-10-27",
				Hour:     "7:00",
				Duration: 90,
			},
		}

		usecase := scheduling.NewAppointmentRescheduler(repo, bus)

		for _, i := range inputs {
			_, err := usecase.Reschedule(i)

			if err == nil {
				t.Errorf("Shoud return an error, got %v", err)
			}

			if !errors.Is(err, scheduling.ErrBusyTime) {
				t.Errorf("The error must be ErrBusyTime, got %v", err)
			}
		}
	})

	t.Run("should_return_error_when_appointment_not_found_in_repository", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:       "1000",
			Date:     "2023-11-31",
			Hour:     "9:15",
			Duration: 120,
		}

		usecase := scheduling.NewAppointmentRescheduler(repo, bus)

		_, err := usecase.Reschedule(i)
		if err == nil {
			t.Errorf("Should return an error, got %v", err)
		}

		if !errors.Is(err, scheduling.ErrAppointmentNotFound) {
			t.Errorf("The error must be ErrAppointmentNotFound, got %v", err)
		}
	})

	t.Run("should_return_error_if_appointment_status_is_different_of_scheduled", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:       "20",
			Date:     "2010-09-18",
			Hour:     "19:15",
			Duration: 60,
		}

		usecase := scheduling.NewAppointmentRescheduler(repo, bus)

		_, err := usecase.Reschedule(i)
		if err == nil {
			t.Errorf("Should return an error, got %v", err)
		}

		if !errors.Is(err, scheduling.ErrInvalidStatusToReschedule) {
			t.Errorf("The error must be ErrInvalidStatusToReschedule, got %v", err)
		}
	})

	t.Run("should_return_an_error_if_the_date_is_in_an_invalid_format", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:       "6",
			Date:     "10-07-2001",
			Hour:     "11:00",
			Duration: 30,
		}
		usecase := scheduling.NewAppointmentRescheduler(repo, bus)

		_, err := usecase.Reschedule(i)
		if err == nil {
			t.Errorf("Shoud return an error, got %v", err)
		}

		if !errors.Is(err, scheduling.ErrInvalidDate) {
			t.Errorf("The error must be ErrInvalidDate, got %v", err)
		}
	})

	t.Run("should_return_an_error_if_the_hour_is_in_an_invalid_format", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:       "2",
			Date:     "2021-07-01",
			Hour:     "11h00",
			Duration: 30,
		}
		usecase := scheduling.NewAppointmentRescheduler(repo, bus)

		_, err := usecase.Reschedule(i)
		if err == nil {
			t.Errorf("Shoud return an error, got %v", err)
		}

		if !errors.Is(err, hour.ErrInvalidHour) {
			t.Errorf("The error must be ErrHourDate, got %v", err)
		}
	})

	t.Run("must_publish_the_rescheduled_appointment_event", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:       "13",
			Date:     "2018-09-15",
			Hour:     "18:00",
			Duration: 30,
		}

		evtPublished := false
		var h event.HandlerFunc = func(e event.Event) {
			evtPublished = true
		}

		bus := event.NewInmemEventBus()
		bus.Subscribe(scheduling.EventAppointmentRescheduled, h)
		usecase := scheduling.NewAppointmentRescheduler(repo, bus)

		_, err := usecase.Reschedule(i)
		if err != nil {
			t.Errorf("Should not return an error, got %v", err)
		}

		if !evtPublished {
			t.Error("The EventAppointmentRescheduled must be published")
		}
	})

	t.Run("must_entry_the_payload_in_reschedule_appointment_event", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:       "12",
			Date:     "2018-05-10",
			Hour:     "11:00",
			Duration: 60,
		}

		evtPublished := false
		var h event.HandlerFunc = func(e event.Event) {
			switch e.Payload().(type) {
			case scheduling.AppointmentReschedulerInput:
				evtPublished = true
			}
		}

		bus := event.NewInmemEventBus()
		bus.Subscribe(scheduling.EventAppointmentRescheduled, h)
		usecase := scheduling.NewAppointmentRescheduler(repo, bus)

		_, err := usecase.Reschedule(i)
		if err != nil {
			t.Errorf("Should not return an error, got %v", err)
		}

		if !evtPublished {
			t.Error("The event payload must be logged")
		}
	})
}
