package session_test

import (
	"errors"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/session"
	"github.com/zafir-co-ao/onna-narciso/internal/session/adapters/inmem"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/id"
)

func TestSessionCloser(t *testing.T) {
	repo := inmem.NewSessionRepository()

	s := session.Session{ID: id.NewID("1")}

	repo.Save(s)

	t.Run("should_close_the_session", func(t *testing.T) {
		sessionID := "1"
		u := session.NewSessionCloser(repo)

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
}
