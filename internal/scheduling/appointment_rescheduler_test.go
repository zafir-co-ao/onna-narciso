package scheduling_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/stubs"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/hour"
)

func TestAppointmentRescheduler(t *testing.T) {
	today := date.Today()

	appointments := []scheduling.Appointment{
		{
			ID:             nanoid.New(),
			ProfessionalID: "1",
			Status:         scheduling.StatusScheduled,
			Date:           date.Date(today.AddDate(0, 1, 5).String()),
			Hour:           "8:00",
			Duration:       240,
		},
		{
			ID:     "20",
			Status: scheduling.StatusCanceled,
		},
		{
			ID:             nanoid.ID("21"),
			ProfessionalID: "2",
			Status:         scheduling.StatusScheduled,
			Date:           today.AddDate(1, 2, 2),
			Hour:           "13:00",
			Duration:       90,
		},
		{
			ID:             nanoid.ID("22"),
			ProfessionalID: "2",
			Status:         scheduling.StatusClosed,
			Date:           today.AddDate(1, 2, 2),
			Hour:           "16:00",
			Duration:       60,
		},
	}

	for i := range 15 {
		a := scheduling.Appointment{ID: nanoid.ID(strconv.Itoa(i)), Status: scheduling.StatusScheduled}
		appointments = append(appointments, a)
	}

	bus := event.NewEventBus()
	sacl := stubs.NewServicesACL()
	pacl := stubs.NewProfessionalsACL()
	repo := scheduling.NewAppointmentRepository(appointments...)
	u := scheduling.NewAppointmentRescheduler(repo, pacl, sacl, bus)

	t.Run("should_reschedule_appointment", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:             "1",
			Date:           today.AddDate(0, 0, 10).String(),
			Hour:           "8:30",
			ProfessionalID: "1",
			ServiceID:      "1",
			Duration:       120,
		}

		err := u.Reschedule(i)
		if !errors.Is(nil, err) {
			t.Errorf("Should not return error, got %v", err)
		}

		a, err := repo.FindByID(nanoid.ID(i.ID))
		if !errors.Is(nil, err) {
			t.Errorf("Should return the appointment not an error, got %v", err)
		}

		if !a.IsRescheduled() {
			t.Errorf("The appointment status must be Rescheduled, got %v", a.Status)
		}
	})

	t.Run("must_enter_the_reschedule_date", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:             "3",
			Date:           today.AddDate(0, 0, 1).String(),
			Hour:           "9:30",
			ProfessionalID: "2",
			ServiceID:      "3",
			Duration:       120,
		}

		err := u.Reschedule(i)
		if !errors.Is(nil, err) {
			t.Errorf("Should not return error, got %v", err)
		}

		o, err := repo.FindByID(nanoid.ID(i.ID))
		if err != nil {
			t.Errorf("Should return the appointment not an error, got %v", err)
		}

		if o.Date.String() != i.Date {
			t.Errorf("The appointment date must be %v, got %v", i.Date, o.Date)
		}
	})

	t.Run("must_enter_the_reschedule_hour", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:             "4",
			Date:           today.AddDate(0, 0, 3).String(),
			Hour:           "10:00",
			ProfessionalID: "3",
			ServiceID:      "4",
			Duration:       120,
		}

		err := u.Reschedule(i)
		if !errors.Is(nil, err) {
			t.Errorf("Should not return error, got %v", err)
		}

		o, err := repo.FindByID(nanoid.ID(i.ID))
		if err != nil {
			t.Errorf("Should return the appointment not an error, got %v", err)
		}

		if o.Hour.String() != i.Hour {
			t.Errorf("The appointment hour must be 10:00, got %v", o.Hour)
		}
	})

	t.Run("must_enter_the_reschedule_duration", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:             "5",
			Date:           today.AddDate(0, 1, 1).String(),
			Hour:           "8:00",
			ProfessionalID: "1",
			ServiceID:      "1",
			Duration:       60,
		}

		err := u.Reschedule(i)
		if !errors.Is(nil, err) {
			t.Errorf("Should not return error, got %v", err)
		}

		o, err := repo.FindByID(nanoid.ID(i.ID))
		if err != nil {
			t.Errorf("Should return the appointment not an error, got %v", err)
		}

		if o.Duration.Value() != i.Duration {
			t.Errorf("The appointment duration must be 60 minutes, got %v", o.Duration)
		}
	})

	t.Run("must_reschedule_an_appointment_more_than_once", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:             "1",
			Date:           today.AddDate(0, 0, 5).String(),
			Hour:           "12:00",
			ProfessionalID: "2",
			ServiceID:      "3",
			Duration:       60,
		}

		err := u.Reschedule(i)
		if !errors.Is(nil, err) {
			t.Errorf("Should not return error, got %v", err)
		}

		o, err := repo.FindByID(nanoid.ID(i.ID))
		if err != nil {
			t.Errorf("Should return the appointment not an error, got %v", err)
		}

		if o.Status != scheduling.StatusRescheduled {
			t.Errorf("The appointment status must be %v, got %v", scheduling.StatusRescheduled, o.Status)
		}
	})

	t.Run("should_return_an_error_if_exists_interception_of_appointments", func(t *testing.T) {
		var inputs = []scheduling.AppointmentReschedulerInput{
			{
				ID:             "8",
				Date:           today.AddDate(0, 1, 5).String(),
				Hour:           "8:00",
				ProfessionalID: "1",
				ServiceID:      "1",
				Duration:       60,
			},
			{
				ID:             "9",
				Date:           today.AddDate(0, 1, 5).String(),
				Hour:           "9:30",
				ProfessionalID: "1",
				ServiceID:      "1",
				Duration:       60,
			},
			{
				ID:             "10",
				Date:           today.AddDate(0, 1, 5).String(),
				Hour:           "11:00",
				ProfessionalID: "1",
				ServiceID:      "1",
				Duration:       60,
			},
			{
				ID:             "10",
				Date:           today.AddDate(0, 1, 5).String(),
				Hour:           "7:00",
				ProfessionalID: "1",
				ServiceID:      "1",
				Duration:       90,
			},
		}

		for _, i := range inputs {
			err := u.Reschedule(i)

			if errors.Is(nil, err) {
				t.Errorf("Shoud return an error, got %v", err)
			}

			if !errors.Is(err, scheduling.ErrBusyTime) {
				t.Errorf("The error must be %v, got %v", scheduling.ErrBusyTime, err)
			}
		}
	})

	t.Run("should_return_error_when_appointment_not_found_in_repository", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:             "1000",
			Date:           "2023-11-31",
			Hour:           "9:15",
			ProfessionalID: "3",
			ServiceID:      "2",
			Duration:       120,
		}

		err := u.Reschedule(i)
		if errors.Is(nil, err) {
			t.Errorf("Should return an error, got %v", err)
		}

		if !errors.Is(err, scheduling.ErrAppointmentNotFound) {
			t.Errorf("The error must be %v, got %v", scheduling.ErrAppointmentNotFound, err)
		}
	})

	t.Run("should_return_error_if_appointment_status_is_different_of_scheduled", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:             "20",
			Date:           today.AddDate(1, 0, 8).String(),
			Hour:           "19:15",
			ProfessionalID: "4",
			ServiceID:      "3",
			Duration:       60,
		}

		err := u.Reschedule(i)
		if errors.Is(nil, err) {
			t.Errorf("Should return an error, got %v", err)
		}

		if !errors.Is(err, scheduling.ErrInvalidStatusToReschedule) {
			t.Errorf("The error must be %v, got %v", scheduling.ErrInvalidStatusToReschedule, err)
		}
	})

	t.Run("should_return_an_error_if_the_date_is_in_an_invalid_format", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:             "6",
			Date:           "10-07-2001",
			Hour:           "11:00",
			ProfessionalID: "2",
			ServiceID:      "3",
			Duration:       30,
		}

		err := u.Reschedule(i)
		if errors.Is(nil, err) {
			t.Errorf("Shoud return an error, got %v", err)
		}

		if !errors.Is(err, date.ErrInvalidFormat) {
			t.Errorf("The error must be %v, got %v", date.ErrInvalidFormat, err)
		}
	})

	t.Run("should_return_an_error_if_the_hour_is_in_an_invalid_format", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:             "2",
			Date:           today.AddDate(1, 0, 8).String(),
			Hour:           "11h00",
			ProfessionalID: "1",
			ServiceID:      "2",
			Duration:       30,
		}
		err := u.Reschedule(i)
		if errors.Is(nil, err) {
			t.Errorf("Shoud return an error, got %v", err)
		}

		if !errors.Is(err, hour.ErrInvalidFormat) {
			t.Errorf("The error must be %v, got %v", hour.ErrInvalidFormat, err)
		}
	})

	t.Run("must_publish_the_rescheduled_appointment_event", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:             "13",
			Date:           today.AddDate(1, 1, 8).String(),
			Hour:           "18:00",
			ProfessionalID: "4",
			ServiceID:      "1",
			Duration:       30,
		}

		evtPublished := false
		var h event.HandlerFunc = func(e event.Event) {
			evtPublished = true
		}

		bus.Subscribe(scheduling.EventAppointmentRescheduled, h)

		err := u.Reschedule(i)
		if !errors.Is(nil, err) {
			t.Errorf("Should not return an error, got %v", err)
		}

		if !evtPublished {
			t.Errorf("The %s must be published", scheduling.EventAppointmentRescheduled)
		}
	})

	t.Run("must_entry_the_payload_in_reschedule_appointment_event", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:             "12",
			Date:           today.AddDate(0, 2, 8).String(),
			Hour:           "11:00",
			ProfessionalID: "3",
			ServiceID:      "4",
			Duration:       60,
		}

		evtPublished := false
		var h event.HandlerFunc = func(e event.Event) {
			switch e.Payload().(type) {
			case scheduling.AppointmentReschedulerInput:
				evtPublished = true
			}
		}

		bus.Subscribe(scheduling.EventAppointmentRescheduled, h)

		err := u.Reschedule(i)
		if !errors.Is(nil, err) {
			t.Errorf("Should not return an error, got %v", err)
		}

		if !evtPublished {
			t.Error("The event payload must be logged")
		}
	})

	t.Run("should_reschedule_the_appointment_with_different_professional", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:             "2",
			Date:           today.AddDate(0, 0, 10).String(),
			ProfessionalID: "1",
			ServiceID:      "2",
			Hour:           "06:00",
			Duration:       60,
		}

		err := u.Reschedule(i)
		if !errors.Is(nil, err) {
			t.Errorf("Should not return an error, got %v", err)
		}

		o, err := repo.FindByID(nanoid.ID(i.ID))
		if !errors.Is(nil, err) {
			t.Errorf("Should return the appointment not an error, got %v", err)
		}

		if o.ProfessionalID.String() != i.ProfessionalID {
			t.Errorf("The Professional ID of appointment must be equal to %s, got %s", i.ProfessionalID, o.ProfessionalID)
		}

		p, err := pacl.FindProfessionalByID(nanoid.ID(i.ProfessionalID))
		if !errors.Is(nil, err) {
			t.Errorf("Should return the professional, got %v", err)
		}

		if o.ProfessionalName != p.Name {
			t.Errorf("The Professional Name of appointment must be equal to %s, got %s", p.Name, o.ServiceName)
		}
	})

	t.Run("should_reschedule_the_appointment_with_different_service", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:             "2",
			Date:           today.AddDate(0, 0, 1).String(),
			ProfessionalID: "2",
			ServiceID:      "3",
			Hour:           "16:00",
			Duration:       30,
		}

		err := u.Reschedule(i)
		if !errors.Is(nil, err) {
			t.Errorf("Should not return an error, got %v", err)
		}

		o, err := repo.FindByID(nanoid.ID(i.ID))
		if !errors.Is(nil, err) {
			t.Errorf("Should return the appointment not an error, got %v", err)
		}

		if o.ServiceID.String() != i.ServiceID {
			t.Errorf("The Service ID of appointment must be equal to %s, got %s", i.ServiceID, o.ServiceID)
		}

		s, err := sacl.FindServiceByID(nanoid.ID(i.ServiceID))
		if !errors.Is(nil, err) {
			t.Errorf("Should return the service, got %v", err)
		}

		if o.ServiceName != s.Name {
			t.Errorf("The Service Name of appointment must be equal to %s, got %s", s.Name, o.ServiceName)
		}
	})

	t.Run("should_return_error_if_professional_not_found_in_acl", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:             "4",
			Date:           "2021-05-04",
			ProfessionalID: "10",
			ServiceID:      "3",
			Hour:           "16:00",
			Duration:       30,
		}

		err := u.Reschedule(i)
		if errors.Is(nil, err) {
			t.Errorf("Shoud return an error, got %v", err)
		}

		if !errors.Is(err, scheduling.ErrProfessionalNotFound) {
			t.Errorf("The error must be %v, got %v", scheduling.ErrProfessionalNotFound, err)
		}
	})

	t.Run("should_return_error_if_service_not_found_in_acl", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:             "7",
			Date:           "2018-05-04",
			ProfessionalID: "1",
			ServiceID:      "10",
			Hour:           "10:00",
			Duration:       90,
		}

		err := u.Reschedule(i)
		if errors.Is(nil, err) {
			t.Errorf("Should return an error, got %v", err)
		}

		if !errors.Is(err, scheduling.ErrServiceNotFound) {
			t.Errorf("The error must be %v, got %v", scheduling.ErrServiceNotFound, err)
		}
	})

	t.Run("should_return_error_if_the_service_is_not_from_the_chosen_professional", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:             "4",
			Date:           "2012-05-04",
			ProfessionalID: "1",
			ServiceID:      "4",
			Hour:           "10:00",
			Duration:       90,
		}

		err := u.Reschedule(i)
		if err == nil {
			t.Errorf("Should return an error, got %v", err)
		}

		if !errors.Is(err, scheduling.ErrInvalidService) {
			t.Errorf("The error must be %v, got %v", scheduling.ErrInvalidService, err)
		}
	})

	t.Run("should_return_error_if_rescheduling_an_appointment_in_the_past", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:             "4",
			Date:           "2014-05-04",
			ProfessionalID: "1",
			ServiceID:      "1",
			Hour:           "10:00",
			Duration:       90,
		}

		err := u.Reschedule(i)
		if errors.Is(nil, err) {
			t.Errorf("Should return an error, got %v", err)
		}

		if !errors.Is(err, date.ErrDateInPast) {
			t.Errorf("The error must be %v, got %v", date.ErrDateInPast, err)
		}
	})

	t.Run("should_return_error_if_rescheduling_in_slot_of_appointment_closed", func(t *testing.T) {
		i := scheduling.AppointmentReschedulerInput{
			ID:             "21",
			Date:           today.AddDate(1, 2, 2).String(),
			ProfessionalID: "2",
			ServiceID:      "3",
			Hour:           "15:00",
			Duration:       120,
		}

		err := u.Reschedule(i)
		if errors.Is(nil, err) {
			t.Errorf("Should return an error, got %v", err)
		}

		if !errors.Is(err, scheduling.ErrBusyTime) {
			t.Errorf("The error must be %v, got %v", scheduling.ErrBusyTime, err)
		}
	})
}
