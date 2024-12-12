package scheduling_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
)

func TestAppointmentCanceler(t *testing.T) {
	appointments := []scheduling.Appointment{
		{ID: "1"},
		{ID: "2", Status: scheduling.StatusCanceled},
		{ID: "3"},
	}

	bus := event.NewEventBus()
	repo := scheduling.NewAppointmentRepository(appointments...)
	u := scheduling.NewAppointmentCanceler(repo, bus)

	t.Run("should_cancel_an_appointment", func(t *testing.T) {
		err := u.Cancel(appointments[0].ID.String())

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		a, err := repo.FindByID(appointments[0].ID)
		if errors.Is(err, scheduling.ErrAppointmentNotFound) {
			t.Errorf("Should return appointment, got %v", err)
		}

		if !a.IsCancelled() {
			t.Errorf("The appointment status should be %v, got: %v", scheduling.StatusCanceled, a.Status)
		}
	})

	t.Run("should_return_error_if_appointment_status_is_canceled", func(t *testing.T) {
		err := u.Cancel(appointments[1].ID.String())

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, scheduling.ErrInvalidStatusToCancel) {
			t.Errorf("The error must be %v, got %v", scheduling.ErrInvalidStatusToClose, err)
		}
	})

	t.Run("should_return_error_if_appointment_not_found", func(t *testing.T) {
		err := u.Cancel("1000")

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, scheduling.ErrAppointmentNotFound) {
			t.Errorf("The error must be %v, got %v", scheduling.ErrAppointmentNotFound, err)
		}
	})

	t.Run("must_publish_the_canceled_appointment_event", func(t *testing.T) {
		a := appointments[2]

		evtAppointmentID := ""
		var h event.HandlerFunc = func(e event.Event) {
			evtAppointmentID = e.Header(event.HeaderAggregateID)
		}

		bus.Subscribe(scheduling.EventAppointmentCanceled, h)

		err := u.Cancel(a.ID.String())
		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if evtAppointmentID == "" {
			t.Errorf("The event must be %v", scheduling.EventAppointmentCanceled)
		}

		if evtAppointmentID != a.ID.String() {
			t.Errorf("The %v must be published with the appointment id, got %v", scheduling.EventAppointmentCanceled, evtAppointmentID)
		}
	})
}
