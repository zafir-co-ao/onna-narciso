package scheduling_test

import (
	"slices"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/hour"
)

func TestWeeklyAppointments(t *testing.T) {
	repo := inmem.NewAppointmentRepository(
		scheduling.Appointment{ID: "1", Date: scheduling.Date("2024-10-09"), Hour: hour.Hour("11:00"), ServiceID: "3"},
		scheduling.Appointment{ID: "4", Date: scheduling.Date("2024-10-25"), Hour: hour.Hour("10:00"), ServiceID: "3"},
		scheduling.Appointment{ID: "5", Date: scheduling.Date("2024-10-11"), Hour: hour.Hour("10:00"), ServiceID: "3"},
		scheduling.Appointment{ID: "6", Date: scheduling.Date("2024-10-12"), Hour: hour.Hour("10:00"), ServiceID: "3"},
		scheduling.Appointment{ID: "7", Date: scheduling.Date("2024-10-22"), Hour: hour.Hour("09:00"), ServiceID: "3"},
	)

	type weeklyAppointmentsTestMatrix struct {
		date          string
		serviceID     string
		professionals []string

		expectedIDs []string
	}

	matrix := []weeklyAppointmentsTestMatrix{
		{date: "2024-10-10", serviceID: "3", professionals: []string{}, expectedIDs: []string{"1", "5", "6"}},
		{date: "2024-10-21", serviceID: "3", professionals: []string{}, expectedIDs: []string{"4", "7"}},
	}

	appointmentsGetter := scheduling.NewWeeklyAppointmentsGetter(repo)

	for _, test := range matrix {

		t.Run(test.date, func(t *testing.T) {

			results, _ := appointmentsGetter.Find(test.date, test.serviceID, []string{})

			if len(results) != len(test.expectedIDs) {
				t.Errorf("Expected %d appointments, got %d", len(test.expectedIDs), len(results))
			}

			for i, appointment := range results {

				if !slices.ContainsFunc(results, func(a scheduling.AppointmentOutput) bool {
					return a.ID == test.expectedIDs[i]
				}) {
					t.Errorf("Expected appointment in IDs %v, got %s", test.expectedIDs, appointment.ID)
				}
			}
		})
	}

}
