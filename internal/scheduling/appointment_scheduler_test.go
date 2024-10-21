package scheduling_test

import (
	"errors"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"
)

func TestAppointmentScheduler(t *testing.T) {
	t.Run("should_schedule_appointment", func(t *testing.T) {
		d := scheduling.AppointmentSchedulerDTO{
			ProfessionalID: "1",
			CustomerID:     "1",
			ServiceID:      "4",
			Date:           "2024-10-09",
			StartHour:      "11:00",
			Duration:       60,
		}
		repo := inmem.NewAppointmentRepository()
		usecase := scheduling.NewAppointmentScheduler(repo)

		id, err := usecase.Schedule(d)

		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if id != "1" {
			t.Error("Appointment id should be 1")
		}
	})

	t.Run("should_store_appointment_in_repository", func(t *testing.T) {
		d := scheduling.AppointmentSchedulerDTO{
			ProfessionalID: "1",
			CustomerID:     "1",
			ServiceID:      "3",
			Date:           "2024-10-09",
			StartHour:      "11:00",
			Duration:       180,
		}
		repo := inmem.NewAppointmentRepository()
		usecase := scheduling.NewAppointmentScheduler(repo)

		id, err := usecase.Schedule(d)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.Get(id)
		if errors.Is(err, scheduling.ErrAppointmentNotFound) {
			t.Error("Appointment should be stored in repository")
		}

		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.ID != id {
			t.Errorf("Appointment ID should be %s", id)
		}
	})

	t.Run("the_status_of_appointment_should_be_scheduled", func(t *testing.T) {
		d := scheduling.AppointmentSchedulerDTO{
			ProfessionalID: "1",
			CustomerID:     "3",
			ServiceID:      "1",
			Date:           "2024-10-09",
			StartHour:      "13:00",
			Duration:       60,
		}
		repo := inmem.NewAppointmentRepository()
		usecase := scheduling.NewAppointmentScheduler(repo)

		id, err := usecase.Schedule(d)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.Get(id)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.IsScheduled() == false {
			t.Errorf("Appointment status should be Scheduled, got %s", appointment.Status)
		}
	})

	t.Run("should_schedule_appointment_with_professional", func(t *testing.T) {
		d := scheduling.AppointmentSchedulerDTO{
			ProfessionalID: "1",
			CustomerID:     "2",
			ServiceID:      "3",
			Date:           "2024-10-12",
			StartHour:      "12:00",
			Duration:       30,
		}
		repo := inmem.NewAppointmentRepository()

		usecase := scheduling.NewAppointmentScheduler(repo)

		id, err := usecase.Schedule(d)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.Get(id)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.ProfessionalID != d.ProfessionalID {
			t.Errorf("The appointment professional must be  %s, got %s", d.ProfessionalID, appointment.ProfessionalID)
		}
	})

	t.Run("should_schedule_appointment_with_customer", func(t *testing.T) {
		d := scheduling.AppointmentSchedulerDTO{
			CustomerID:     "1",
			ProfessionalID: "2",
			ServiceID:      "4",
			Date:           "2024-10-12",
			StartHour:      "10:00",
			Duration:       30,
		}
		repo := inmem.NewAppointmentRepository()

		usecase := scheduling.NewAppointmentScheduler(repo)

		id, err := usecase.Schedule(d)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.Get(id)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.CustomerID != d.CustomerID {
			t.Errorf("The appointment customer must be  %s, got %s", d.CustomerID, appointment.CustomerID)
		}
	})

	t.Run("should_schedule_appointmet_with_service", func(t *testing.T) {
		d := scheduling.AppointmentSchedulerDTO{
			ProfessionalID: "3",
			ServiceID:      "10",
			CustomerID:     "2",
			Date:           "2024-09-17",
			StartHour:      "10:30",
			Duration:       30,
		}
		repo := inmem.NewAppointmentRepository()

		usecase := scheduling.NewAppointmentScheduler(repo)

		id, err := usecase.Schedule(d)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.Get(id)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.ServiceID != d.ServiceID {
			t.Errorf("The appointment service must be  %s, got %s", d.ServiceID, appointment.ServiceID)
		}
	})

	t.Run("should_schedule_appointment_with_date", func(t *testing.T) {
		d := scheduling.AppointmentSchedulerDTO{
			ProfessionalID: "3",
			ServiceID:      "4",
			CustomerID:     "2",
			Date:           "2024-05-05",
			StartHour:      "18:00",
			Duration:       60,
		}
		repo := inmem.NewAppointmentRepository()

		usecase := scheduling.NewAppointmentScheduler(repo)

		id, err := usecase.Schedule(d)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.Get(id)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.Date != d.Date {
			t.Errorf("The appointment date must be %s, got %s", d.Date, appointment.Date)
		}
	})

	t.Run("should_schedule_appointment_with_start_hour", func(t *testing.T) {
		d := scheduling.AppointmentSchedulerDTO{
			ProfessionalID: "2",
			ServiceID:      "4",
			CustomerID:     "2",
			Date:           "2024-11-10",
			StartHour:      "11:30",
			Duration:       120,
		}
		repo := inmem.NewAppointmentRepository()

		usecase := scheduling.NewAppointmentScheduler(repo)

		id, err := usecase.Schedule(d)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.Get(id)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.Start != d.StartHour {
			t.Errorf("The appointment start hour must be %s, got %s", d.StartHour, appointment.Start)
		}
	})

	t.Run("should_schedule_appointment_with_duration", func(t *testing.T) {
		d := scheduling.AppointmentSchedulerDTO{
			ProfessionalID: "3",
			CustomerID:     "3",
			ServiceID:      "4",
			Date:           "2024-10-15",
			StartHour:      "8:00",
			Duration:       30,
		}
		repo := inmem.NewAppointmentRepository()

		usecase := scheduling.NewAppointmentScheduler(repo)

		id, err := usecase.Schedule(d)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.Get(id)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.Duration != d.Duration {
			t.Errorf("The appointment duration must be 30, got %d", appointment.Duration)
		}
	})

	t.Run("must_calculate_the_end_time_of_the_appointment", func(t *testing.T) {
		d := scheduling.AppointmentSchedulerDTO{
			ProfessionalID: "3",
			CustomerID:     "3",
			ServiceID:      "4",
			Date:           "2024-10-15",
			StartHour:      "8:35",
			Duration:       90,
		}
		repo := inmem.NewAppointmentRepository()

		usecase := scheduling.NewAppointmentScheduler(repo)

		id, err := usecase.Schedule(d)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.Get(id)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.End != "10:05" {
			t.Errorf("The appointment end hour must be 10:05, got %s", appointment.End)
		}
	})

	t.Run("should_return_the_busy_time_error_when_there_is_not_availability_in_schedule", func(t *testing.T) {
		d := scheduling.AppointmentSchedulerDTO{
			ProfessionalID: "3",
			CustomerID:     "3",
			ServiceID:      "4",
			Date:           "2024-10-15",
			StartHour:      "8:00",
			Duration:       90,
		}

		a, _ := scheduling.NewAppointment("1", "3", "4", "3", "2024-10-15", "8:00", 90)
		repo := inmem.NewAppointmentRepository()
		repo.Save(a)

		usecase := scheduling.NewAppointmentScheduler(repo)

		_, err := usecase.Schedule(d)
		if err == nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if !errors.Is(err, scheduling.ErrBusyTime) {
			t.Errorf("The error must be ErrBusyTime, got %v", err)
		}
	})
}
