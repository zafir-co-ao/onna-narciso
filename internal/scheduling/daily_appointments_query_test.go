package scheduling_test

import (
	"slices"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"
)

func TestDailyAppointments(t *testing.T) {

	repo := inmem.NewAppointmentRepository()

	repo.Save(scheduling.Appointment{
		ID:    "1",
		Date:  "2024-10-10",
		Start: "11:00",
	})

	repo.Save(scheduling.Appointment{
		ID:    "4",
		Date:  "2024-10-10",
		Start: "10:00",
	})

	repo.Save(scheduling.Appointment{
		ID:    "5",
		Date:  "2024-10-11",
		Start: "10:00",
	})

	repo.Save(scheduling.Appointment{
		ID:    "6",
		Date:  "2024-10-12",
		Start: "10:00",
	})

	repo.Save(scheduling.Appointment{
		ID:    "7",
		Date:  "2024-10-12",
		Start: "09:00",
	})

	type dailyAppointmentsTestMatrix struct {
		date        string
		expectedIDs []string
	}

	matrix := []dailyAppointmentsTestMatrix{
		{date: "2024-10-10", expectedIDs: []string{"1", "4"}},
		{date: "2024-10-11", expectedIDs: []string{"5"}},
		{date: "2024-10-12", expectedIDs: []string{"6", "7"}},
	}

	appointmentsGetter := scheduling.NewDayliAppointmentsGetter(repo)

	for _, test := range matrix {

		t.Run(test.date, func(t *testing.T) {

			results, _ := appointmentsGetter.Get(test.date)

			if len(results) != len(test.expectedIDs) {
				t.Errorf("Expected %d appointments, got %d", len(test.expectedIDs), len(results))
			}

			for i, appointment := range results {

				if !slices.ContainsFunc(results, func(a scheduling.Appointment) bool {
					return a.ID == test.expectedIDs[i]
				}) {
					t.Errorf("Expected appointment in IDs %v, got %s", test.expectedIDs, appointment.ID)
				}

			}
		})
	}

}
