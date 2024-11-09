package session_test

import (
	"errors"
	"testing"
	"time"

	"github.com/zafir-co-ao/onna-narciso/internal/session"
	"github.com/zafir-co-ao/onna-narciso/internal/session/adapters/inmem"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/event"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/id"
)

func TestSessionCloser(t *testing.T) {
	bus := event.NewInmemEventBus()
	repo := inmem.NewSessionRepository()
	u := session.NewSessionCloser(repo, bus)

	s1 := session.Session{ID: id.NewID("1")}
	s2 := session.Session{ID: id.NewID("2")}
	s3 := session.Session{ID: id.NewID("3"), Status: session.StatusClosed}
	s4 := session.Session{ID: id.NewID("4")}

	repo.Save(s1)
	repo.Save(s2)
	repo.Save(s3)
	repo.Save(s4)

	t.Run("should_close_the_session", func(t *testing.T) {
		sessionID := "1"

		err := u.Close(sessionID)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(id.NewID(sessionID))
		if !errors.Is(nil, err) {
			t.Errorf("Should return the session in repository, got %v", err)
		}

		if !s.IsClosed() {
			t.Errorf("The session should be closed, got %v", s.Status)
		}
	})

	t.Run("must_record_the_closing_time_of_the_session", func(t *testing.T) {
		sessionID := "2"

		err := u.Close(sessionID)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(id.NewID(sessionID))
		if !errors.Is(nil, err) {
			t.Errorf("Should return the session in repository, got %v", err)
		}

		if s.CloseTime.Hour() != time.Now().Hour() {
			t.Errorf("The session close hour should be equal with hour in clock, got %v", s.CloseTime.Hour())
		}
	})

	t.Run("must_publish_the_session_closed_event", func(t *testing.T) {
		var IsPublished bool = false

		var h event.HandlerFunc = func(e event.Event) {
			IsPublished = true
		}

		bus.Subscribe(session.EventSessionClosed, h)

		err := u.Close("4")

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if !IsPublished {
			t.Errorf("The EventSessionClosed must be publised, got %v", IsPublished)
		}
	})

	t.Run("should_return_error_if_session_not_exists_in_repository", func(t *testing.T) {
		sessionID := "200"

		err := u.Close(sessionID)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(session.ErrSessionNotFound, err) {
			t.Errorf("The error must be ErrSessionNotFound, got %v", err)
		}
	})

	t.Run("should_return_error_if_the_session_is_already_closed", func(t *testing.T) {
		sessionID := "3"

		err := u.Close(sessionID)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(session.ErrSessionClosed, err) {
			t.Errorf("The error must be ErrSessionClosed, got %v", err)
		}
	})
}
