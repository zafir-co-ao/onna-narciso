package sessions_test

import (
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/sessions"
)

func TestSessionFinder(t *testing.T) {
	s := []sessions.Session{
		{
			ID:            "1",
			AppointmentID: "1",
		},
		{
			ID:            "2",
			AppointmentID: "2",
		},
		{
			ID:            "3",
			AppointmentID: "3",
		},
	}

	repo := sessions.NewInmemRepository(s...)

	type sessionFinderTestMatrix struct {
		appointmentIDs []string
		expectedIDs    []string
	}

	matrix := []sessionFinderTestMatrix{
		{appointmentIDs: []string{"1", "3"}, expectedIDs: []string{"1", "3"}},
		{appointmentIDs: []string{"2", "3"}, expectedIDs: []string{"2", "3"}},
		{appointmentIDs: []string{"1", "4"}, expectedIDs: []string{"1"}},
	}

	finder := sessions.NewSessionFinder(repo)

	for _, test := range matrix {
		t.Run("should_find_sessions_with_appointments_ids", func(t *testing.T) {

			results, _ := finder.Find(test.appointmentIDs)

			if len(results) != len(test.expectedIDs) {
				t.Errorf("Expected %d sessions, got %d", len(test.expectedIDs), len(results))
			}

			for i, session := range results {
				if session.ID != test.expectedIDs[i] {
					t.Errorf("Expected Session ID: %s, got %s", session.ID, test.expectedIDs[i])
				}
			}
		})
	}
}
