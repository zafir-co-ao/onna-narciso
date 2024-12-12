package scheduling_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
)

func TestAppointmentCanceler(t *testing.T) {
	repo := scheduling.NewAppointmentRepository()
	bus := event.NewEventBus()
	u := scheduling.NewAppointmentCanceler(repo, bus)

	a1 := scheduling.Appointment{ID: "1"}
	a2 := scheduling.Appointment{ID: "2", Status: scheduling.StatusCanceled}
	a3 := scheduling.Appointment{ID: "3"}
	a4 := scheduling.Appointment{ID: "4"}

	_ = repo.Save(a1)
	_ = repo.Save(a2)
	_ = repo.Save(a3)
	_ = repo.Save(a4)

	t.Run("should_cancel_an_appointment", func(t *testing.T) {
		err := u.Cancel("1")
		if err != nil {
			t.Errorf("Canceling appointment should not return error: %v", err)
		}

		a, err := repo.FindByID("1")
		if errors.Is(err, scheduling.ErrAppointmentNotFound) {
			t.Errorf("Appointment should be stored in repository")
		}

		if !a.IsCancelled() {
			t.Errorf("Appointment should be status cancelled, got: %v", a.Status)
		}
	})

	t.Run("should_return_error_if_appointment_status_is_canceled", func(t *testing.T) {
		err := u.Cancel("2")
		if err == nil {
			t.Errorf("Canceling appointment should return error: %v", err)
		}

		if !errors.Is(err, scheduling.ErrInvalidStatusToCancel) {
			t.Errorf("Canceling appointment should return error: %v", err)
		}
	})

	t.Run("should_return_error_appointment_not_found_when_appointment_not_exists_in_repository", func(t *testing.T) {
		err := u.Cancel("1000")
		if err == nil {
			t.Errorf("Canceling appointment should return error: %v", err)
		}

		if !errors.Is(err, scheduling.ErrAppointmentNotFound) {
			t.Errorf("Canceling appointment should return error: %v", err)
		}
	})

	t.Run("must_publish_the_canceled_appointment_event", func(t *testing.T) {
		id := "3"

		evtAppointmentID := ""
		var h event.HandlerFunc = func(e event.Event) {
			evtAppointmentID = e.Header(event.HeaderAggregateID)
		}

		bus.Subscribe(scheduling.EventAppointmentCanceled, h)

		err := u.Cancel(id)
		if err != nil {
			t.Errorf("Should not return an error, got %v", err)
		}

		if evtAppointmentID == "" {
			t.Error("The EventAppointmentCanceled must be published")
		}

		if evtAppointmentID != id {
			t.Errorf("The EventAppointmentCanceled must be published with the appointment id, got %v", evtAppointmentID)
		}
	})
}
