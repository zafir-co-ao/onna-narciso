package scheduling_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/adapters/inmem"

	"github.com/zafir-co-ao/onna-narciso/internal/scheduling/stubs"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/hour"
)

func TestAppointmentScheduler(t *testing.T) {
	bus := event.NewEventBus()
	repo := inmem.NewAppointmentRepository()
	a1 := scheduling.Appointment{ID: "1", Date: "2024-10-14", Hour: "8:00", Duration: 90, Status: scheduling.StatusScheduled, ProfessionalID: nanoid.ID("3")}
	a2 := scheduling.Appointment{ID: "2", Date: "2024-10-15", Hour: "8:00", Duration: 480, Status: scheduling.StatusScheduled, ProfessionalID: nanoid.ID("2")}
	a3 := scheduling.Appointment{ID: "6", Date: "2020-04-01", Hour: "19:00", Duration: 60, Status: scheduling.StatusCanceled}

	repo.Save(a1)
	repo.Save(a2)
	repo.Save(a3)

	cacl := stubs.NewCustomersACL()
	pacl := stubs.NewProfessionalsACL()
	sacl := stubs.NewServicesACL()

	t.Run("should_schedule_appointment", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "1",
			CustomerID:     "1",
			ServiceID:      "4",
			Date:           "2024-10-09",
			Hour:           "11:00",
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
			Hour:           "11:00",
			Duration:       180,
		}
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.FindByID(nanoid.ID(o.ID))
		if errors.Is(err, scheduling.ErrAppointmentNotFound) {
			t.Error("Appointment should be stored in repository")
		}

		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.ID.String() != o.ID {
			t.Errorf("Appointment ID should be %s", o.ID)
		}
	})

	t.Run("the_status_of_appointment_should_be_scheduled", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "1",
			CustomerID:     "3",
			ServiceID:      "1",
			Date:           "2024-10-09",
			Hour:           "13:00",
			Duration:       60,
		}
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.FindByID(nanoid.ID(o.ID))
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
			Hour:           "12:00",
			Duration:       30,
		}
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.FindByID(nanoid.ID(o.ID))
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.ProfessionalID.String() != i.ProfessionalID {
			t.Errorf("The appointment professional must be  %s, got %s", i.ProfessionalID, appointment.ProfessionalID)
		}
	})

	t.Run("should_schedule_appointment_with_customer", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			CustomerID:     "1",
			ProfessionalID: "2",
			ServiceID:      "4",
			Date:           "2024-10-12",
			Hour:           "10:00",
			Duration:       30,
		}
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.FindByID(nanoid.ID(o.ID))
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.CustomerID.String() != i.CustomerID {
			t.Errorf("The appointment customer must be  %s, got %s", i.CustomerID, appointment.CustomerID)
		}
	})

	t.Run("should_schedule_appointmet_with_service", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			ServiceID:      "1",
			CustomerID:     "2",
			Date:           "2024-09-17",
			Hour:           "10:30",
			Duration:       30,
		}
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.FindByID(nanoid.ID(o.ID))
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.ServiceID.String() != i.ServiceID {
			t.Errorf("The appointment service must be  %s, got %s", i.ServiceID, appointment.ServiceID)
		}
	})

	t.Run("should_schedule_appointment_with_date", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			ServiceID:      "4",
			CustomerID:     "2",
			Date:           "2024-05-05",
			Hour:           "18:00",
			Duration:       60,
		}
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.FindByID(nanoid.ID(o.ID))
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.Date.String() != i.Date {
			t.Errorf("The appointment date must be %s, got %s", i.Date, appointment.Date)
		}
	})

	t.Run("should_schedule_appointment_with_start_hour", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "2",
			ServiceID:      "4",
			CustomerID:     "2",
			Date:           "2024-11-10",
			Hour:           "11:30",
			Duration:       120,
		}
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.FindByID(nanoid.ID(o.ID))
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.Hour.String() != i.Hour {
			t.Errorf("The appointment start hour must be %s, got %s", i.Hour, appointment.Hour)
		}
	})

	t.Run("should_schedule_appointment_with_duration", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerID:     "3",
			ServiceID:      "4",
			Date:           "2024-10-25",
			Hour:           "8:00",
			Duration:       30,
		}
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		appointment, err := repo.FindByID(nanoid.ID(o.ID))
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		if appointment.Duration != i.Duration {
			t.Errorf("The appointment duration must be 30, got %d", appointment.Duration)
		}
	})

	t.Run("should_return_the_busy_time_error_when_there_is_not_availability_in_schedule", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerID:     "3",
			ServiceID:      "4",
			Date:           "2024-10-14",
			Hour:           "8:00",
			Duration:       90,
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		_, err := usecase.Schedule(i)
		if !errors.Is(err, scheduling.ErrBusyTime) {
			t.Errorf("The error must be ErrBusyTime, got %v", err)
		}
	})

	t.Run("should_return_the_busy_time_error_when_the_appointment_clashes_with_one_on_schedule", func(t *testing.T) {
		var inputs = []scheduling.AppointmentSchedulerInput{
			{ProfessionalID: "2", CustomerID: "3", ServiceID: "4", Date: "2024-10-15", Hour: "9:00", Duration: 30},
			{ProfessionalID: "2", CustomerID: "3", ServiceID: "4", Date: "2024-10-15", Hour: "9:30", Duration: 60},
			{ProfessionalID: "2", CustomerID: "3", ServiceID: "4", Date: "2024-10-15", Hour: "7:30", Duration: 60},
			{ProfessionalID: "2", CustomerID: "2", ServiceID: "4", Date: "2024-10-15", Hour: "11:30", Duration: 480},
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		for _, i := range inputs {
			_, err := usecase.Schedule(i)

			if !errors.Is(err, scheduling.ErrBusyTime) {
				t.Errorf("The error must be ErrBusyTime, got %v", err)
			}
		}
	})

	t.Run("should_return_error_customer_not_found_if_customer_not_exists", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerID:     "4",
			ServiceID:      "4",
			Date:           "2024-10-15",
			Hour:           "8:00",
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

	t.Run("should_return_error_professional_not_found_if_professional_not_exists", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "400",
			CustomerID:     "100",
			ServiceID:      "4",
			Date:           "2024-09-01",
			Hour:           "8:00",
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

	t.Run("should_return_error_service_not_found_if_service_not_exists", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerID:     "3",
			ServiceID:      "5",
			Date:           "2024-09-01",
			Hour:           "8:00",
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
			Hour:           "8:00",
			Duration:       60,
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		a, _ := repo.FindByID(nanoid.ID(o.ID))

		customer, err := cacl.FindCustomerByID(a.CustomerID)
		if err != nil {
			t.Errorf("Should return customer: %v", err)
		}

		if customer.ID != a.CustomerID {
			t.Errorf("The customer must be the same as the appointment")
		}
	})

	t.Run("should_return_error_when_not_register_a_customer", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			ServiceID:      "4",
			Date:           "2024-08-01",
			Hour:           "8:00",
			Duration:       60,
		}

		scheduler := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		_, err := scheduler.Schedule(i)
		if err == nil {
			t.Errorf("Scheduling appointment should return error: %v", err)
		}

		if !errors.Is(err, scheduling.ErrCustomerNotFound) {
			t.Errorf("The error must be ErrCustomerNotFound, got %v", err)
		}
	})

	t.Run("should_return_error_if_date_of_appointment_is_in_wrong_format", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerID:     "3",
			ServiceID:      "4",
			Date:           "01/08/2024",
			Hour:           "8:00",
			Duration:       60,
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		_, err := usecase.Schedule(i)
		if err == nil {
			t.Errorf("Scheduling appointment should return error: %v", err)
		}

		if !errors.Is(err, date.ErrInvalidDate) {
			t.Errorf("The error must be ErrInvalidDate, got %v", err)
		}
	})

	t.Run("should_return_error_if_start_hour_of_appointment_is_in_wrong_format", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerID:     "3",
			ServiceID:      "4",
			Date:           "2024-08-01",
			Hour:           "8h00",
			Duration:       60,
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		_, err := usecase.Schedule(i)
		if err == nil {
			t.Errorf("Scheduling appointment should return error: %v", err)
		}

		if !errors.Is(err, hour.ErrInvalidHour) {
			t.Errorf("The error must be ErrInvalidHour, got %v", err)
		}
	})

	t.Run("should_generate_the_id_of_appointment", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerID:     "3",
			ServiceID:      "4",
			Date:           "2024-08-01",
			Hour:           "8:00",
			Duration:       60,
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		a, err := repo.FindByID(nanoid.ID(o.ID))
		if err != nil {
			t.Errorf("Should return appointment: %v", err)
		}

		if a.ID.String() != o.ID {
			t.Errorf("The ID of appointment must be the same as the generated")
		}
	})

	t.Run("must_register_the_name_of_professional_on_the_appointment", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "3",
			CustomerID:     "3",
			ServiceID:      "4",
			Date:           "2022-07-01",
			Hour:           "8:00",
			Duration:       60,
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		a, err := repo.FindByID(nanoid.ID(o.ID))
		if err != nil {
			t.Errorf("Should return appointment: %v", err)
		}

		p, _ := pacl.FindProfessionalByID(a.ProfessionalID)
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
			Hour:           "19:00",
			Duration:       60,
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		a, err := repo.FindByID(nanoid.ID(o.ID))
		if err != nil {
			t.Errorf("Should return appointment: %v", err)
		}

		c, _ := cacl.FindCustomerByID(a.CustomerID)
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
			Hour:           "19:00",
			Duration:       60,
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		a, err := repo.FindByID(nanoid.ID(o.ID))
		if err != nil {
			t.Errorf("Should return appointment: %v", err)
		}

		s, _ := sacl.FindServiceByID(a.ServiceID)
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
			Hour:           "19:00",
			Duration:       60,
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Scheduling appointment should not return error: %v", err)
		}

		a, err := repo.FindByID(nanoid.ID(o.ID))
		if err != nil {
			t.Errorf("Should return appointment: %v", err)
		}

		if a.ID.String() != o.ID {
			t.Errorf("The ID of appointment must be the same as the generated: %v", a.ID.String())
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
			Hour:           "9:00",
			Duration:       60,
		}

		evtPublished := false
		var h event.HandlerFunc = func(e event.Event) {
			evtPublished = true
		}

		bus.Subscribe(scheduling.EventAppointmentScheduled, h)
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		_, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Should not return an error, got %v", err)
		}

		if !evtPublished {
			t.Error("The EventAppointmentCanceled must be published")
		}

	})

	t.Run("must_entry_the_payload_in_schedule_appointment_event", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "1",
			CustomerID:     "2",
			ServiceID:      "1",
			Date:           "2020-02-01",
			Hour:           "10:00",
			Duration:       60,
		}

		evtAggID := ""
		var h event.HandlerFunc = func(e event.Event) {
			switch e.Payload().(type) {
			case scheduling.AppointmentSchedulerInput:
				evtAggID = e.Header(event.HeaderAggregateID)
			}
		}

		bus.Subscribe(scheduling.EventAppointmentScheduled, h)
		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		o, err := usecase.Schedule(i)
		if err != nil {
			t.Errorf("Should not return an error, got %v", err)
		}

		if evtAggID != o.ID {

			t.Errorf("Event header Aggregate ID should equal Output ID, got: %v != %v", evtAggID, o.ID)
		}
	})

	t.Run("should_return_error_when_customer_name_nor_customer_phone_provided", func(t *testing.T) {
		i := scheduling.AppointmentSchedulerInput{
			ProfessionalID: "1",
			CustomerID:     "",
			CustomerName:   "",
			CustomerPhone:  "",
			ServiceID:      "1",
			Date:           "2024-02-11",
			Hour:           "15:00",
			Duration:       60,
		}

		usecase := scheduling.NewAppointmentScheduler(repo, cacl, pacl, sacl, bus)

		_, err := usecase.Schedule(i)

		if err == nil {
			t.Errorf("Should return an error, got %v", err)
		}

		if !errors.Is(err, scheduling.ErrCustomerNotFound) {
			t.Errorf("Should return an ErrorCustomerNotFound, got %v", err)
		}
	})
}
