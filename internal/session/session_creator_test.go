package session_test

import (
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/session"
	"github.com/zafir-co-ao/onna-narciso/internal/session/adapters/inmem"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/id"
	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
)

func TestSessionCreator(t *testing.T) {
	t.Run("should_create_session", func(t *testing.T) {
		repo := inmem.NewSessionRepository()

		creator := session.NewSessionCreator(repo)
		aid := testdata.Appointments[0].ID.String()

		_, err := creator.Create(aid)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("should_store_session_in_the_repository", func(t *testing.T) {
		repo := inmem.NewSessionRepository()
		creator := session.NewSessionCreator(repo)
		aid := testdata.Appointments[1].ID.String()

		session, err := creator.Create(aid)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(id.NewID(session.ID))
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if s.ID.String() != session.ID {
			t.Errorf("Expected the session to be persisted, got %v", s)
		}

		if s.AppointmentID.String() != aid {
			t.Errorf("Expected the session to have the appointment ID, got %v", s.AppointmentID)
		}

	})

}
