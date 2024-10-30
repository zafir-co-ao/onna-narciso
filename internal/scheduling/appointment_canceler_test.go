package scheduling_test

import (
	"errors"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/event"
)

func TestAppointmentCanceler(t *testing.T) {
	repo := inmem.NewAppointmentRepository()
	bus := event.NewInmemEventBus()

	a1 := scheduling.Appointment{ID: "1"}
	a2 := scheduling.Appointment{ID: "2", Status: scheduling.StatusCanceled}
	a3 := scheduling.Appointment{ID: "3"}
	a4 := scheduling.Appointment{ID: "4"}

	repo.Save(a1)
	repo.Save(a2)
	repo.Save(a3)
	repo.Save(a4)

	t.Run("should_cancel_an_appointment", func(t *testing.T) {
		usecase := scheduling.NewAppointmentCanceler(repo, bus)

		err := usecase.Execute("1")
		if err != nil {
			t.Errorf("Canceling appointment should not return error: %v", err)
		}

		app, err := repo.FindByID("1")
		if errors.Is(scheduling.ErrAppointmentNotFound, err) {
			t.Errorf("Appointment should be stored in repository")
		}

		if !app.IsCancelled() {
			t.Errorf("Appointment should be status cancelled, got: %v", app.Status)
		}
	})

	t.Run("should_return_error_if_appointment_status_is_canceled", func(t *testing.T) {
		usecase := scheduling.NewAppointmentCanceler(repo, bus)

		err := usecase.Execute("2")
		if err == nil {
			t.Errorf("Canceling appointment should return error: %v", err)
		}

		if !errors.Is(scheduling.ErrInvalidStatusToCancel, err) {
			t.Errorf("Canceling appointment should return error: %v", err)
		}
	})

	t.Run("should_return_error_appointment_not_found_when_appointment_not_exists_in_repository", func(t *testing.T) {
		usecase := scheduling.NewAppointmentCanceler(repo, bus)

		err := usecase.Execute("1000")
		if err == nil {
			t.Errorf("Canceling appointment should return error: %v", err)
		}

		if !errors.Is(scheduling.ErrAppointmentNotFound, err) {
			t.Errorf("Canceling appointment should return error: %v", err)
		}
	})

	t.Run("must_publish_the_canceled_appointment_event", func(t *testing.T) {
		id := "3"
		h := &FakeStorageHandler{}
		bus := event.NewInmemEventBus()
		bus.Subscribe(scheduling.EventAppointmentCanceled, h)

		usecase := scheduling.NewAppointmentCanceler(repo, bus)

		err := usecase.Execute(id)
		if !errors.Is(nil, err) {
			t.Errorf("Should not return an error, got %v", err)
		}

		if !h.WasPublished(id, scheduling.EventAppointmentCanceled) {
			t.Error("The EventAppointmentCanceled must be published")
		}
	})

	t.Run("must_entry_the_appointment_id_when_publish_event", func(t *testing.T) {
		id := "4"
		h := &FakeStorageHandler{}
		bus := event.NewInmemEventBus()
		bus.Subscribe(scheduling.EventAppointmentCanceled, h)

		usecase := scheduling.NewAppointmentCanceler(repo, bus)

		err := usecase.Execute(id)
		if !errors.Is(nil, err) {
			t.Errorf("Should not return an error, got %v", err)
		}

		if !h.WasPublished(id, scheduling.EventAppointmentCanceled) {
			t.Error("The EventAppointmentCanceled must be published")
		}
	})
}

type FakeStorageHandler struct {
	events []event.Event
}

func (s *FakeStorageHandler) Handle(e event.Event) {
	s.events = append(s.events, e)
}

func (s *FakeStorageHandler) WasPublished(id, name string) bool {
	for _, e := range s.events {
		if e.Name() == name && e.Header(event.HeaderAggregateID) == id {
			return true
		}
	}
	return false
}

func (s *FakeStorageHandler) FindEventByAggregateID(id string) (event.Event, error) {
	for _, e := range s.events {
		if e.Header(event.HeaderAggregateID) == id {
			return e, nil
		}
	}
	return event.Event{}, event.ErrEventNotFound
}
