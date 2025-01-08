package sessions_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions/stubs"
	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
)

func TestSessionCreator(t *testing.T) {
	bus := event.NewEventBus()
	aacl := stubs.NewSchedulingServiceACL()
	repo := sessions.NewInmemRepository()
	u := sessions.NewSessionCreator(repo, aacl, bus)

	t.Run("should_create_session", func(t *testing.T) {
		a := testdata.Appointments[0]

		_, err := u.Create(a.ID.String())

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("should_store_session_in_the_repository", func(t *testing.T) {
		a := testdata.Appointments[1]

		o, err := u.Create(a.ID.String())
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(nanoid.ID(o.ID))
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if s.ID.String() != o.ID {
			t.Errorf("Expected the session to be persisted, got %v", s)
		}

		if s.AppointmentID.String() != a.ID.String() {
			t.Errorf("Expected the session to have the appointment ID, got %v", s.AppointmentID)
		}
	})

	t.Run("should_publish_SessionCheckedIn_event", func(t *testing.T) {
		a := testdata.Appointments[2]

		epublished := false
		var h event.HandlerFunc = func(e event.Event) {
			epublished = true
		}

		bus.Subscribe(sessions.EventSessionCheckedIn, h)

		_, err := u.Create(a.ID.String())
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if !epublished {
			t.Errorf("Expected the event to be published")
		}
	})

	t.Run("should_fill_session_with_appointment_client_professional_and_service", func(t *testing.T) {
		a := testdata.Appointments[2]

		o, err := u.Create(a.ID.String())
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(nanoid.ID(o.ID))
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if s.AppointmentID != a.ID {
			t.Errorf("Expected the session to have the appointment ID, got %v", s.AppointmentID.String())
		}

		if s.CustomerID != a.CustomerID {
			t.Errorf("Expected the session to have the client ID, got %v", s.CustomerID.String())
		}

		if s.Services[0].ProfessionalID != a.ProfessionalID {
			t.Errorf("Expected the session to have the professional ID, got %v", s.Services[0].ProfessionalID.String())
		}

		if s.Services[0].ID != a.ServiceID {
			t.Errorf("Expected the session to have the service ID, got %v", s.Services[0].ID.String())
		}
	})

	t.Run("should_return_error_if_the_appointment_has_already_been_canceled", func(t *testing.T) {
		a := testdata.Appointments[4]

		_, err := u.Create(a.ID.String())
		if err == nil {
			t.Errorf("Expected an error, got %v", err)
		}

		if !errors.Is(err, sessions.ErrAppointmentCanceled) {
			t.Errorf("The error must be %v, got %v", sessions.ErrAppointmentCanceled, err)
		}
	})

	t.Run("should_return_error_if_appointment_alrealy_closed", func(t *testing.T) {
		a := testdata.Appointments[6]

		_, err := u.Create(a.ID.String())
		if errors.Is(nil, err) {
			t.Errorf("Expected an error, got %v", err)
		}

		if !errors.Is(err, sessions.ErrAppointmentClosed) {
			t.Errorf("The error must be %v, got %v", sessions.ErrAppointmentClosed, err)
		}
	})

	t.Run("should_return_error_if_the_checkin_date_is_different_from_the_appointment_date", func(t *testing.T) {
		a := testdata.Appointments[5]

		_, err := u.Create(a.ID.String())
		if errors.Is(nil, err) {
			t.Errorf("Expected an error, got %v", err)
		}

		if !errors.Is(err, sessions.ErrInvalidCheckinDate) {
			t.Errorf("The error must be %v, got %v", sessions.ErrInvalidCheckinDate, err)
		}
	})
}
