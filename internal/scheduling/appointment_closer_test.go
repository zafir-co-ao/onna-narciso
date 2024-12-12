package scheduling_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
)

func TestAppointmentCloser(t *testing.T) {
	appointments := []scheduling.Appointment{
		{ID: "1"},
		{ID: "2", Status: scheduling.StatusClosed},
		{ID: "3"},
	}

	bus := event.NewEventBus()
	repo := scheduling.NewAppointmentRepository(appointments...)
	u := scheduling.NewAppointmentCloser(repo, bus)

	t.Run("should_close_appointment", func(t *testing.T) {
		id := appointments[0].ID

		err := u.Close(id.String())
		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		a, err := repo.FindByID(id)
		if !errors.Is(nil, err) {
			t.Errorf("Should return the appointment in repository, got %v", err)
		}

		if !a.IsClosed() {
			t.Errorf("The appointment status should be %v, got %v", scheduling.StatusClosed, a.Status)
		}
	})

	t.Run("must_publish_the_closing_appointment_event", func(t *testing.T) {
		id := appointments[2].ID.String()

		isPublished := false
		var h event.HandlerFunc = func(e event.Event) {
			if e.Header(event.HeaderAggregateID) == id {
				isPublished = true
			}
		}

		bus.Subscribe(scheduling.EventAppointmentClosed, h)

		err := u.Close(id)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if !isPublished {
			t.Errorf("The %v must be published", scheduling.EventAppointmentClosed)
		}
	})

	t.Run("should_return_error_if_appointment_not_found", func(t *testing.T) {
		err := u.Close(nanoid.New().String())

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, scheduling.ErrAppointmentNotFound) {
			t.Errorf("The error mus be %v, got %v", scheduling.ErrAppointmentNotFound, err)
		}
	})

	t.Run("should_return_error_if_appointment_already_closed", func(t *testing.T) {
		id := appointments[1].ID.String()

		err := u.Close(id)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, scheduling.ErrInvalidStatusToClose) {
			t.Errorf("The error must be %v, got %v", scheduling.ErrInvalidStatusToClose, err)
		}
	})
}
