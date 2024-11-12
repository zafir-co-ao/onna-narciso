package session_test

import (
	"errors"
	"testing"
	"time"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/session"
	"github.com/zafir-co-ao/onna-narciso/internal/session/adapters/inmem"
)

func TestSessionStarter(t *testing.T) {
	repo := inmem.NewSessionRepository()
	bus := event.NewEventBus()

	s1 := session.Session{ID: "1"}
	s2 := session.Session{ID: "2", Status: session.StatusStarted}
	s3 := session.Session{ID: "3"}
	s4 := session.Session{ID: "4"}

	_ = repo.Save(s1)
	_ = repo.Save(s2)
	_ = repo.Save(s3)
	_ = repo.Save(s4)

	u := session.NewSessionStarter(repo, bus)

	t.Run("should_start_session", func(t *testing.T) {
		id := "1"

		err := u.Start(id)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(nanoid.ID(id))
		if !errors.Is(nil, err) {
			t.Errorf("Should return the session in repository, got %v", err)
		}

		if !s.IsStarted() {
			t.Errorf("The session should be started, got %v", s.Status)
		}
	})

	t.Run("must_record_the_starting_time_of_the_session", func(t *testing.T) {
		id := "3"
		err := u.Start("3")

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(nanoid.ID(id))
		if !errors.Is(nil, err) {
			t.Errorf("Should return the session in repository, got %v", err)
		}

		if s.StartTime.Hour() != time.Now().Hour() {
			t.Errorf("The session start hour should be equal with hour in clock, got %v", s.StartTime.Hour())
		}
	})

	t.Run("must_publish_the_session_started_event", func(t *testing.T) {
		id := "4"

		evtAggID := ""
		var isPublished bool = false
		var h event.HandlerFunc = func(e event.Event) {
			evtAggID = e.Header(event.HeaderAggregateID)
			isPublished = true
		}

		bus.Subscribe(session.EventSessionStarted, h)

		err := u.Start(id)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if !isPublished {
			t.Errorf("The EventSessionStarted must be publised, got %v", isPublished)
		}

		if evtAggID != id {
			t.Errorf("Event header Aggregate ID should equal ID: %v, got: %v", id, evtAggID)
		}

	})

	t.Run("should_return_error_if_session_not_found_in_repository", func(t *testing.T) {

		err := u.Start("100")

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(session.ErrSessionNotFound, err) {
			t.Errorf("The error must be ErrSessionNotFound, got %v", err)
		}
	})

	t.Run("should_return_error_if_the_session_is_already_started", func(t *testing.T) {

		err := u.Start("2")

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(session.ErrSessionStarted, err) {
			t.Errorf("The error must be ErrSessionStarted, got %v", err)
		}
	})
}
