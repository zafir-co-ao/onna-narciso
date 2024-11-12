package session_test

import (
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/session"
	"github.com/zafir-co-ao/onna-narciso/internal/session/adapters/inmem"
	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
)

func TestSessionCreator(t *testing.T) {

	bus := event.NewEventBus()

	t.Run("should_create_session", func(t *testing.T) {
		repo := inmem.NewSessionRepository()

		creator := session.NewSessionCreator(repo, bus)
		aid := testdata.Appointments[0].ID.String()

		_, err := creator.Create(aid)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

	})

	t.Run("should_store_session_in_the_repository", func(t *testing.T) {
		repo := inmem.NewSessionRepository()
		creator := session.NewSessionCreator(repo, event.NewEventBus())
		aid := testdata.Appointments[1].ID.String()

		session, err := creator.Create(aid)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(nanoid.ID(session.ID))
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

	t.Run("should_publish_SessionCheckedIn_event", func(t *testing.T) {
		b := event.NewEventBus()
		c := session.NewSessionCreator(inmem.NewSessionRepository(), b)
		aid := testdata.Appointments[2].ID.String()

		epublished := false
		var h event.HandlerFunc = func(e event.Event) {
			epublished = true
		}

		b.Subscribe("SessionCheckedIn", h)

		_, err := c.Create(aid)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if !epublished {
			t.Errorf("Expected the event to be published")
		}
	})
}
