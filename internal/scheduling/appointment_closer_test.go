package scheduling_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
)

func TestAppointmentCloser(t *testing.T) {
	bus := event.NewEventBus()
	repo := scheduling.NewAppointmentRepository()
	u := scheduling.NewAppointmentCloser(repo, bus)

	a1 := scheduling.Appointment{ID: "1"}
	a2 := scheduling.Appointment{ID: "2", Status: scheduling.StatusClosed}
	a3 := scheduling.Appointment{ID: "3"}

	_ = repo.Save(a1)
	_ = repo.Save(a2)
	_ = repo.Save(a3)

	t.Run("should_close_appointment", func(t *testing.T) {
		id := "1"

		err := u.Close(id)
		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		a, err := repo.FindByID(nanoid.ID(id))
		if !errors.Is(nil, err) {
			t.Errorf("Should return the appointment in repository, got %v", err)
		}

		if !a.IsClosed() {
			t.Errorf("The appointment should be closed, got %v", a.Status)
		}
	})

	t.Run("must_publish_the_closing_appointment_event", func(t *testing.T) {
		id := "3"

		evtAggID := ""
		isPublished := false
		var h event.HandlerFunc = func(e event.Event) {
			evtAggID = e.Header(event.HeaderAggregateID)
			isPublished = true
		}

		bus.Subscribe(scheduling.EventAppointmentClosed, h)

		err := u.Close(id)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if !isPublished {
			t.Error("The EventAppointmentClosed must be published")
		}

		if evtAggID != id {
			t.Errorf("The EventAppointmentClosed must be published with the appointment id, got %v", evtAggID)
		}
	})

	t.Run("should_return_error_if_appointment_not_exists_in_repository", func(t *testing.T) {
		err := u.Close(nanoid.New().String())

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(scheduling.ErrAppointmentNotFound, err) {
			t.Errorf("Closing appointment should return error: %v", err)
		}
	})

	t.Run("should_return_error_if_appointment_already_closed", func(t *testing.T) {
		err := u.Close("2")

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(scheduling.ErrInvalidStatusToClose, err) {
			t.Errorf("Closing appointment should return error: %v", err)
		}
	})
}
