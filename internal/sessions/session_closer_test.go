package sessions_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions/stubs"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/hour"
)

func TestSessionCloser(t *testing.T) {
	_sessions := []sessions.Session{
		{ID: "20", Status: sessions.StatusCheckedIn},
	}

	for i := range 10 {
		s := sessions.Session{ID: nanoid.ID(strconv.Itoa(i)), Status: sessions.StatusStarted}
		_sessions = append(_sessions, s)
	}

	bus := event.NewEventBus()
	repo := sessions.NewInmemRepository(_sessions...)
	sacl := stubs.NewServicesServiceACL()

	u := sessions.NewSessionCloser(repo, sacl, bus)

	t.Run("should_close_the_session", func(t *testing.T) {
		i := sessions.CloserInput{
			SessionID:   "1",
			ServicesIDs: []string{"1", "2"},
		}

		err := u.Close(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(nanoid.ID(i.SessionID))
		if !errors.Is(nil, err) {
			t.Errorf("Should return the session in repository, got %v", err)
		}

		if !s.IsClosed() {
			t.Errorf("The session should be closed, got %v", s.Status)
		}
	})

	t.Run("must_record_the_closing_time_of_the_session", func(t *testing.T) {
		i := sessions.CloserInput{
			SessionID:   "2",
			ServicesIDs: []string{"2", "3"},
		}

		err := u.Close(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(nanoid.ID(i.SessionID))
		if !errors.Is(nil, err) {
			t.Errorf("Should return the session in repository, got %v", err)
		}

		if s.CloseTime != hour.Now() {
			t.Errorf("The session close hour should be equal with hour in clock, got %v", s.CloseTime)
		}
	})

	t.Run("must_entry_the_additional_services_in_session", func(t *testing.T) {
		i := sessions.CloserInput{
			SessionID:   "5",
			ServicesIDs: []string{"1", "2"},
		}

		err := u.Close(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(nanoid.ID(i.SessionID))

		if !errors.Is(nil, err) {
			t.Errorf("Should return the session in repository, got %v", err)
		}

		if s.Services[0].ServiceID.String() != "1" {
			t.Errorf("The Service ID must be 1, got %s", s.Services[0].ServiceID.String())
		}

		if s.Services[1].ServiceID.String() != "2" {
			t.Errorf("The Service ID must be 2, got %s", s.Services[1].ServiceID.String())
		}
	})

	t.Run("must_entry_the_professional_in_session", func(t *testing.T) {
		i := sessions.CloserInput{
			SessionID:   "6",
			ServicesIDs: []string{"2", "3"},
		}

		err := u.Close(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(nanoid.ID(i.SessionID))
		if !errors.Is(nil, err) {
			t.Errorf("Should return the session in repository, got %v", err)
		}

		if s.Services[0].ProfessionalID.String() != "1" {
			t.Errorf("The Professional ID must be 1, got %v", s.Services[0].ProfessionalID.String())
		}

		if s.Services[1].ProfessionalID.String() != "1" {
			t.Errorf("The Professional ID must be 1, got %v", s.Services[1].ProfessionalID.String())
		}
	})

	t.Run("should_close_the_session_without_additional_services", func(t *testing.T) {
		i := sessions.CloserInput{
			SessionID:   "8",
			ServicesIDs: []string{},
		}

		err := u.Close(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(nanoid.ID(i.SessionID))
		if !errors.Is(nil, err) {
			t.Errorf("Should return the session in repository, got %v", err)
		}

		if len(s.Services) != 0 {
			t.Errorf("The session must be closed without services, got %v", len(s.Services))
		}
	})

	t.Run("should_return_error_if_not_found_service_acl", func(t *testing.T) {
		i := sessions.CloserInput{
			SessionID:   "7",
			ServicesIDs: []string{"1", "2", "10"},
		}

		err := u.Close(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, sessions.ErrServiceNotFound) {
			t.Errorf("The error must be %v, got %v", sessions.ErrServiceNotFound, err)
		}
	})

	t.Run("must_publish_the_session_closed_event", func(t *testing.T) {
		i := sessions.CloserInput{
			SessionID:   "4",
			ServicesIDs: make([]string, 0),
		}

		evtAggID := ""
		var isPublished bool = false
		var h event.HandlerFunc = func(e event.Event) {
			switch e.Payload().(type) {
			case sessions.CloserInput:
				evtAggID = e.Header(event.HeaderAggregateID)
				isPublished = true
			}
		}

		bus.Subscribe(sessions.EventSessionClosed, h)

		err := u.Close(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if !isPublished {
			t.Errorf("The EventSessionClosed must be publised, got %v", isPublished)
		}

		if evtAggID != i.SessionID {
			t.Errorf("Event header Aggregate ID should equal ID: %v, got: %v", i.SessionID, evtAggID)
		}
	})

	t.Run("should_return_error_if_session_not_exists_in_repository", func(t *testing.T) {
		i := sessions.CloserInput{
			SessionID:   "200",
			ServicesIDs: make([]string, 0),
		}

		err := u.Close(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, sessions.ErrSessionNotFound) {
			t.Errorf("The error must be %v, got %v", sessions.ErrServiceNotFound, err)
		}
	})

	t.Run("should_return_error_if_the_session_is_already_closed", func(t *testing.T) {
		i := sessions.CloserInput{
			SessionID:   "3",
			ServicesIDs: make([]string, 0),
		}

		u.Close(i)

		err := u.Close(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, sessions.ErrSessionClosed) {
			t.Errorf("The error must be %v, got %v", sessions.ErrSessionClosed, err)
		}
	})

	t.Run("should_return_error_if_session_status_not_is_started", func(t *testing.T) {
		i := sessions.CloserInput{
			SessionID:   "20",
			ServicesIDs: make([]string, 0),
		}

		err := u.Close(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, sessions.ErrInvalidStatusToClose) {
			t.Errorf("The error must be %v, got %v", sessions.ErrInvalidStatusToClose, err)
		}
	})
}
