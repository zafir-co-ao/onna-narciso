package sessions_test

import (
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions/adapters/inmem"
	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
)

func TestSessionCreator(t *testing.T) {

	var aacl sessions.AppointmentsACLFunc = func(id nanoid.ID) (sessions.Appointment, error) {
		for _, a := range testdata.Appointments {
			if a.ID == id {
				return sessions.Appointment{
					ID:               a.ID,
					ProfessionalID:   a.ProfessionalID,
					CustomerID:       a.CustomerID,
					ServiceID:        a.ServiceID,
					CustomerName:     string(a.CustomerName),
					ProfessionalName: string(a.ProfessionalName),
					ServiceName:      string(a.ServiceName),
				}, nil
			}
		}

		return sessions.Appointment{}, sessions.ErrAppointmentNotFound
	}

	bus := event.NewEventBus()

	t.Run("should_create_session", func(t *testing.T) {
		repo := inmem.NewSessionRepository()

		creator := sessions.NewSessionCreator(repo, bus, aacl)
		aid := testdata.Appointments[0].ID.String()

		_, err := creator.Create(aid)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

	})

	t.Run("should_store_session_in_the_repository", func(t *testing.T) {
		repo := inmem.NewSessionRepository()
		creator := sessions.NewSessionCreator(repo, event.NewEventBus(), aacl)
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
		c := sessions.NewSessionCreator(inmem.NewSessionRepository(), b, aacl)
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

	t.Run("should_fill_session_with_appointment_client_professional_and_service", func(t *testing.T) {
		repo := inmem.NewSessionRepository()
		creator := sessions.NewSessionCreator(repo, event.NewEventBus(), aacl)
		appointment := testdata.Appointments[2]

		out, err := creator.Create(appointment.ID.String())
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(nanoid.ID(out.ID))
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if s.AppointmentID != appointment.ID {
			t.Errorf("Expected the session to have the appointment ID, got %v", s.AppointmentID.String())
		}

		if s.CustomerID != appointment.CustomerID {
			t.Errorf("Expected the session to have the client ID, got %v", s.CustomerID.String())
		}

		if s.Services[0].ProfessionalID != appointment.ProfessionalID {
			t.Errorf("Expected the session to have the professional ID, got %v", s.Services[0].ProfessionalID.String())
		}

		if s.Services[0].ServiceID != appointment.ServiceID {
			t.Errorf("Expected the session to have the service ID, got %v", s.Services[0].ServiceID.String())
		}
	})
}
