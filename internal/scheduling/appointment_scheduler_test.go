package scheduling_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/tests/stubs"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/event"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/id"
)

func TestAppointmentScheduler(t *testing.T) {
	bus := event.NewInmemEventBus()
	repo := inmem.NewAppointmentRepository()
	a1 := scheduling.Appointment{ID: "1", Date: "2024-10-14", Start: "8:00", Duration: 90, Status: scheduling.StatusScheduled}
	a2 := scheduling.Appointment{ID: "2", Date: "2024-10-15", Start: "8:00", Duration: 480, Status: scheduling.StatusScheduled}
	a3 := scheduling.Appointment{ID: "6", Date: "2020-04-01", Start: "19:00", Duration: 60, Status: scheduling.StatusCanceled}

	repo.Save(a1)
	repo.Save(a2)
	repo.Save(a3)

	cacl := stubs.CustomerAclStub{}
	pacl := stubs.Pacl
	sacl := stubs.Sacl

	t.Run("should_schedule_appointment", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "1",
			CustomerID:     "1",
			ServiceID:      "4",
			Date:           "2024-10-09",
			StartHour:      "11:00",
			Duration:       60,
		}
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if o.ID == "" {
			t.Errorf("Appointment id should be %v", o.ID)
		}
	})

	t.Run("should_store_appointment_in_repository", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "1",
			CustomerID:     "1",
			ServiceID:      "3",
			Date:           "2024-09-09",
			StartHour:      "11:00",
			Duration:       180,
		}
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.FindByID(id.NewID(o.ID))
		if errors.Is(err, scheduling.ErrAppointmentNotFound) {
			t.Error("Appointment should be stored in repository")
		}

		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.ID.Value() != o.ID {
			t.Errorf("Appointment ID should be %s", o.ID)
		}
	})

	t.Run("the_status_of_appointment_should_be_scheduled", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "1",
			CustomerID:     "3",
			ServiceID:      "1",
			Date:           "2024-10-09",
			StartHour:      "13:00",
			Duration:       60,
		}
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.FindByID(id.NewID(o.ID))
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.IsScheduled() == false {
			t.Errorf("Appointment status should be Scheduled, got %s", appointment.Status)
		}
	})

	t.Run("should_schedule_appointment_with_professional", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "1",
			CustomerID:     "2",
			ServiceID:      "3",
			Date:           "2024-10-12",
			StartHour:      "12:00",
			Duration:       30,
		}
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.FindByID(id.NewID(o.ID))
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.ProfessionalID.Value() != i.ProfessionalID {
			t.Errorf("The appointment professional must be  %s, got %s", i.ProfessionalID, appointment.ProfessionalID)
		}
	})

	t.Run("should_schedule_appointment_with_customer", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			CustomerID:     "1",
			ProfessionalID: "2",
			ServiceID:      "4",
			Date:           "2024-10-12",
			StartHour:      "10:00",
			Duration:       30,
		}
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.FindByID(id.NewID(o.ID))
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.CustomerID.Value() != i.CustomerID {
			t.Errorf("The appointment customer must be  %s, got %s", i.CustomerID, appointment.CustomerID)
		}
	})

	t.Run("should_schedule_appointmet_with_service", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			ServiceID:      "1",
			CustomerID:     "2",
			Date:           "2024-09-17",
			StartHour:      "10:30",
			Duration:       30,
		}
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.FindByID(id.NewID(o.ID))
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.ServiceID.Value() != i.ServiceID {
			t.Errorf("The appointment service must be  %s, got %s", i.ServiceID, appointment.ServiceID)
		}
	})

	t.Run("should_schedule_appointment_with_date", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			ServiceID:      "4",
			CustomerID:     "2",
			Date:           "2024-05-05",
			StartHour:      "18:00",
			Duration:       60,
		}
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.FindByID(id.NewID(o.ID))
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.Date.Value() != i.Date {
			t.Errorf("The appointment date must be %s, got %s", i.Date, appointment.Date)
		}
	})

	t.Run("should_schedule_appointment_with_start_hour", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "2",
			ServiceID:      "4",
			CustomerID:     "2",
			Date:           "2024-11-10",
			StartHour:      "11:30",
			Duration:       120,
		}
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.FindByID(id.NewID(o.ID))
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.Start.Value() != i.StartHour {
			t.Errorf("The appointment start hour must be %s, got %s", i.StartHour, appointment.Start)
		}
	})

	t.Run("should_schedule_appointment_with_duration", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerID:     "3",
			ServiceID:      "4",
			Date:           "2024-10-25",
			StartHour:      "8:00",
			Duration:       30,
		}
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.FindByID(id.NewID(o.ID))
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.Duration != i.Duration {
			t.Errorf("The appointment duration must be 30, got %d", appointment.Duration)
		}
	})

	t.Run("must_calculate_the_end_time_of_the_appointment", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerID:     "3",
			ServiceID:      "4",
			Date:           "2024-05-15",
			StartHour:      "8:35",
			Duration:       90,
		}
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		a, err := repo.FindByID(id.NewID(o.ID))
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if a.End != "10:05" {
			t.Errorf("The appointment end hour must be 10:05, got %s", a.End)
		}
	})

	t.Run("should_return_the_busy_time_error_when_there_is_not_availability_in_schedule", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerID:     "3",
			ServiceID:      "4",
			Date:           "2024-10-14",
			StartHour:      "8:00",
			Duration:       90,
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		_, err := usecase.Schedule(i)
		if err == nil {
			t.Errorf("Scheduling appointment should return error: %v", err)
		}

		if !errors.Is(err, scheduling.ErrBusyTime) {
			t.Errorf("The error must be ErrBusyTime, got %v", err)
		}
	})

	t.Run("should_return_the_busy_time_error_when_the_appointment_clashes_with_one_on_schedule", func(t *testing.T) {
		var inputs = []scheduling.AppointmentSchedulerInput{
			{
				ProfessionalID: "1",
				CustomerID:     "3",
				ServiceID:      "4",
				Date:           "2024-10-15",
				StartHour:      "9:00",
				Duration:       30,
			},
			{
				ProfessionalID: "1",
				CustomerID:     "3",
				ServiceID:      "4",
				Date:           "2024-10-15",
				StartHour:      "9:30",
				Duration:       60,
			},
			{
				ProfessionalID: "3",
				CustomerID:     "3",
				ServiceID:      "4",
				Date:           "2024-10-15",
				StartHour:      "7:30",
				Duration:       60,
			},
			{
				ProfessionalID: "2",
				CustomerID:     "2",
				ServiceID:      "4",
				Date:           "2024-10-15",
				StartHour:      "11:30",
				Duration:       480,
			},
		}
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		for _, i := range inputs {
			_, err := usecase.Schedule(i)
			if err == nil {
				t.Errorf("Scheduling appointment should return error: %v", err)
			}
			if !errors.Is(err, scheduling.ErrBusyTime) {
				t.Errorf("The error must be ErrBusyTime, got %v", err)
			}
		}
	})

	t.Run("should_return_error_customer_not_found_if_customer_not_exists_in_repository", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerID:     "4",
			ServiceID:      "4",
			Date:           "2024-10-15",
			StartHour:      "8:00",
			Duration:       90,
		}
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		_, err := usecase.Schedule(i)
		if err == nil {
			t.Errorf("Scheduling appointment should return error: %v", err)
		}

		if !errors.Is(err, scheduling.ErrCustomerNotFound) {
			t.Errorf("The error must be ErrCustomerNotFound, got %v", err)
		}
	})

	t.Run("should_return_error_professional_not_found_if_professional_not_exists_in_repository", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "4",
			CustomerID:     "100",
			ServiceID:      "4",
			Date:           "2024-09-01",
			StartHour:      "8:00",
			Duration:       60,
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		_, err := usecase.Schedule(i)
		if err == nil {
			t.Errorf("Scheduling appointment should return error: %v", err)
		}

		if !errors.Is(err, scheduling.ErrProfessionalNotFound) {
			t.Errorf("The error must be ErrProfessionalNotFound, got %v", err)
		}
	})

	t.Run("should_return_error_service_not_found_if_service_not_exists_in_repository", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerID:     "3",
			ServiceID:      "5",
			Date:           "2024-09-01",
			StartHour:      "8:00",
			Duration:       60,
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		_, err := usecase.Schedule(i)
		if err == nil {
			t.Errorf("Scheduling appointment should return error: %v", err)
		}

		if !errors.Is(err, scheduling.ErrServiceNotFound) {
			t.Errorf("The error must be ErrServiceNotFound, got %v", err)
		}
	})

	t.Run("must_register_the_customer_at_the_time_of_the_appointment", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerName:   "John Doe",
			CustomerPhone:  "123456789",
			ServiceID:      "4",
			Date:           "2024-09-01",
			StartHour:      "8:00",
			Duration:       60,
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		a, _ := repo.FindByID(id.NewID(o.ID))

		customer, err := cacl.FindCustomerByID(a.CustomerID.Value())
		if err != nil {
			t.Errorf("Should return customer: %v", err)
		}

		if customer.ID.Value() != a.CustomerID.Value() {
			t.Errorf("The customer must be the same as the appointment")
		}
	})

	t.Run("should_return_error_when_not_register_a_customer", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			ServiceID:      "4",
			Date:           "2024-08-01",
			StartHour:      "8:00",
			Duration:       60,
		}

		scheduler := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		_, err := scheduler.Schedule(i)
		if err == nil {
			t.Errorf("Scheduling appointment should return error: %v", err)
		}

		if !errors.Is(err, scheduling.ErrCustomerRegistration) {
			t.Errorf("The error must be ErrCustomerRegistration, got %v", err)
		}
	})

	t.Run("should_return_error_if_date_of_appointment_is_in_wrong_format", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerID:     "3",
			ServiceID:      "4",
			Date:           "01/08/2024",
			StartHour:      "8:00",
			Duration:       60,
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		_, err := usecase.Schedule(i)
		if err == nil {
			t.Errorf("Scheduling appointment should return error: %v", err)
		}

		if !errors.Is(err, scheduling.ErrInvalidDate) {
			t.Errorf("The error must be ErrInvalidDate, got %v", err)
		}
	})

	t.Run("should_return_error_if_start_hour_of_appointment_is_in_wrong_format", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerID:     "3",
			ServiceID:      "4",
			Date:           "2024-08-01",
			StartHour:      "8h00",
			Duration:       60,
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		_, err := usecase.Schedule(i)
		if err == nil {
			t.Errorf("Scheduling appointment should return error: %v", err)
		}

		if !errors.Is(err, scheduling.ErrInvalidHour) {
			t.Errorf("The error must be ErrInvalidHour, got %v", err)
		}
	})

	t.Run("should_generate_the_id_of_appointment", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerID:     "3",
			ServiceID:      "4",
			Date:           "2024-08-01",
			StartHour:      "8:00",
			Duration:       60,
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		a, err := repo.FindByID(id.NewID(o.ID))
		if err != nil {
			t.Errorf("Should return appointment: %v", err)
		}

		if a.ID.Value() != o.ID {
			t.Errorf("The ID of appointment must be the same as the generated")
		}
	})

	t.Run("must_register_the_name_of_professional_on_the_appointment", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerID:     "3",
			ServiceID:      "4",
			Date:           "2022-07-01",
			StartHour:      "8:00",
			Duration:       60,
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		a, err := repo.FindByID(id.NewID(o.ID))
		if err != nil {
			t.Errorf("Should return appointment: %v", err)
		}

		p, _ := pacl.FindProfessionalByID(a.ProfessionalID.Value())
		if a.ProfessionalName != p.Name {
			t.Errorf("The professional name must be the same as the appointment: %v", a.ProfessionalName)
		}
	})

	t.Run("must_register_the_customer_name_in_the_appointment", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerID:     "2",
			ServiceID:      "4",
			Date:           "2024-07-01",
			StartHour:      "19:00",
			Duration:       60,
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		a, err := repo.FindByID(id.NewID(o.ID))
		if err != nil {
			t.Errorf("Should return appointment: %v", err)
		}

		c, _ := cacl.FindCustomerByID(a.CustomerID.Value())
		if a.CustomerName != c.Name {
			t.Errorf("The customer name must be the same as the appointment: %v", a.CustomerName)
		}
	})

	t.Run("must_register_the_service_name_in_the_appointment", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerID:     "2",
			ServiceID:      "4",
			Date:           "2024-05-01",
			StartHour:      "19:00",
			Duration:       60,
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		a, err := repo.FindByID(id.NewID(o.ID))
		if err != nil {
			t.Errorf("Should return appointment: %v", err)
		}

		s, _ := sacl.FindServiceByID(a.ServiceID.Value())
		if a.ServiceName != s.Name {
			t.Errorf("The service name must be the same as the appointment: %v", a.ServiceName)
		}
	})

	t.Run("must_schedule_an_appointment_at_the_time_of_canceled_appointment", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerID:     "2",
			ServiceID:      "4",
			Date:           "2020-04-01",
			StartHour:      "19:00",
			Duration:       60,
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		a, err := repo.FindByID(id.NewID(o.ID))
		if err != nil {
			t.Errorf("Should return appointment: %v", err)
		}

		if a.ID.Value() != o.ID {
			t.Errorf("The ID of appointment must be the same as the generated: %v", a.ID.Value())
		}

		if a.Status != scheduling.StatusScheduled {
			t.Errorf("The status of appointment must be scheduled: %v", a.Status)
		}
	})

	t.Run("must_publish_the_scheduled_appointment_event", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerID:     "2",
			ServiceID:      "1",
			Date:           "2020-01-01",
			StartHour:      "9:00",
			Duration:       60,
		}
		h := &FakeStorageHandler{}
		bus.Subscribe(scheduling.EventAppointmentScheduled, h)
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if !errors.Is(nil, err) {
			t.Errorf("Should not return an error, got %v", err)
		}

		if !h.WasPublished(o.ID, scheduling.EventAppointmentScheduled) {
			t.Error("The EventAppointmentScheduled must be published")
		}
	})

	t.Run("must_entry_the_payload_in_schedule_appointment_event", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "1",
			CustomerID:     "2",
			ServiceID:      "1",
			Date:           "2020-02-01",
			StartHour:      "10:00",
			Duration:       60,
		}
		h := &FakeStorageHandler{}
		bus.Subscribe(scheduling.EventAppointmentScheduled, h)
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if !errors.Is(nil, err) {
			t.Errorf("Should not return an error, got %v", err)
		}

		e, err := h.FindEventByAggregateID(o.ID)
		if errors.Is(event.ErrEventNotFound, err) {
			t.Errorf("Should return an event, got %v", err)

		}

		if reflect.TypeOf(e.Payload()) != reflect.TypeOf(scheduling.AppointmentSchedulerInput{}) {
			t.Error("The event payload must be logged")
		}
	})
}
