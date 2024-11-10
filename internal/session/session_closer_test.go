package session_test

import (
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/session"
	"github.com/zafir-co-ao/onna-narciso/internal/session/adapters/inmem"
)

func TestSessionCloser(t *testing.T) {
	bus := event.NewEventBus()
	repo := inmem.NewSessionRepository()
	u := session.NewSessionCloser(repo, sacl, bus)

	for i := range 10 {
		s := session.Session{ID: nanoid.ID(strconv.Itoa(i + 1))}
		repo.Save(s)
	}

	t.Run("should_close_the_session", func(t *testing.T) {
		input := session.SessionCloserInput{
			SessionID:   "1",
			ServicesIDs: []string{"1", "2"},
		}

		err := u.Close(input)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(nanoid.ID(input.SessionID))
		if !errors.Is(nil, err) {
			t.Errorf("Should return the session in repository, got %v", err)
		}

		if !s.IsClosed() {
			t.Errorf("The session should be closed, got %v", s.Status)
		}
	})

	t.Run("must_record_the_closing_time_of_the_session", func(t *testing.T) {
		input := session.SessionCloserInput{
			SessionID:   "2",
			ServicesIDs: []string{"2", "3"},
		}

		err := u.Close(input)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(nanoid.ID(input.SessionID))
		if !errors.Is(nil, err) {
			t.Errorf("Should return the session in repository, got %v", err)
		}

		if s.CloseTime.Hour() != time.Now().Hour() {
			t.Errorf("The session close hour should be equal with hour in clock, got %v", s.CloseTime.Hour())
		}
	})

	t.Run("must_entry_the_additional_services_in_session", func(t *testing.T) {
		input := session.SessionCloserInput{
			SessionID:   "5",
			ServicesIDs: []string{"1", "2"},
		}

		err := u.Close(input)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(nanoid.ID(input.SessionID))

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
		input := session.SessionCloserInput{
			SessionID:   "6",
			ServicesIDs: []string{"2", "3"},
		}

		err := u.Close(input)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(nanoid.ID(input.SessionID))
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
		input := session.SessionCloserInput{
			SessionID:   "8",
			ServicesIDs: []string{},
		}

		err := u.Close(input)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(nanoid.ID(input.SessionID))
		if !errors.Is(nil, err) {
			t.Errorf("Should return the session in repository, got %v", err)
		}

		if len(s.Services) != 0 {
			t.Errorf("The session must be closed without services, got %v", len(s.Services))
		}
	})

	t.Run("should_return_error_if_not_found_service_acl", func(t *testing.T) {
		input := session.SessionCloserInput{
			SessionID:   "7",
			ServicesIDs: []string{"1", "2", "10"},
		}

		err := u.Close(input)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(session.ErrServiceNotFound, err) {
			t.Errorf("The error must be ErrServiceNotFound, got %v", err)
		}
	})

	t.Run("must_publish_the_session_closed_event", func(t *testing.T) {
		input := session.SessionCloserInput{
			SessionID:   "4",
			ServicesIDs: make([]string, 0),
		}

		evtAggID := ""
		var isPublished bool = false
		var h event.HandlerFunc = func(e event.Event) {
			switch e.Payload().(type) {
			case session.SessionCloserInput:
				evtAggID = e.Header(event.HeaderAggregateID)
				isPublished = true
			}
		}

		bus.Subscribe(session.EventSessionClosed, h)

		err := u.Close(input)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if !isPublished {
			t.Errorf("The EventSessionClosed must be publised, got %v", isPublished)
		}

		if evtAggID != input.SessionID {
			t.Errorf("Event header Aggregate ID should equal ID: %v, got: %v", input.SessionID, evtAggID)
		}
	})

	t.Run("should_return_error_if_session_not_exists_in_repository", func(t *testing.T) {
		input := session.SessionCloserInput{
			SessionID:   "200",
			ServicesIDs: make([]string, 0),
		}

		err := u.Close(input)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(session.ErrSessionNotFound, err) {
			t.Errorf("The error must be ErrSessionNotFound, got %v", err)
		}
	})

	t.Run("should_return_error_if_the_session_is_already_closed", func(t *testing.T) {
		input := session.SessionCloserInput{
			SessionID:   "3",
			ServicesIDs: make([]string, 0),
		}

		u.Close(input)

		err := u.Close(input)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(session.ErrSessionClosed, err) {
			t.Errorf("The error must be ErrSessionClosed, got %v", err)
		}
	})
}

var services = map[string]session.Service{
	"1": session.Service{ServiceID: nanoid.ID("1"), ProfessionalID: nanoid.ID("1")},
	"2": session.Service{ServiceID: nanoid.ID("2"), ProfessionalID: nanoid.ID("1")},
	"3": session.Service{ServiceID: nanoid.ID("3"), ProfessionalID: nanoid.ID("1")},
}

var sacl session.ServiceACLFunc = func(ids []nanoid.ID) ([]session.Service, error) {
	selectedServices := make([]session.Service, 0)
	for _, v := range ids {
		s := services[v.String()]
		if s.ServiceID.String() != v.String() {
			return []session.Service{}, session.ErrServiceNotFound
		}

		selectedServices = append(selectedServices, s)
	}

	return selectedServices, nil
}
