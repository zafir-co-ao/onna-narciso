package scheduling

import (
	"github.com/zafir-co-ao/onna-narciso/internal/shared/event"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/id"
)

const EventAppointmentScheduled = "EventAppointmentScheduled"

type AppointmentSchedulerInput struct {
	ID             string
	ProfessionalID string
	CustomerID     string
	CustomerName   string
	CustomerPhone  string
	ServiceID      string
	Date           string
	StartHour      string
	Duration       int
}

type AppointmentScheduler interface {
	Schedule(i AppointmentSchedulerInput) (AppointmentOutput, error)
}

type appointmentScedulerImpl struct {
	repo            AppointmentRepository
	serviceAcl      ServiceAcl
	customerAcl     CustomerAcl
	professionalAcl ProfessionalAcl
	bus             event.Bus
}

func NewAppointmentScheduler(
	repo AppointmentRepository,
	cacl CustomerAcl,
	pacl ProfessionalAcl,
	sacl ServiceAcl,
	bus event.Bus,
) AppointmentScheduler {
	return &appointmentScedulerImpl{
		repo:            repo,
		customerAcl:     cacl,
		professionalAcl: pacl,
		serviceAcl:      sacl,
		bus:             bus,
	}
}

func (u *appointmentScedulerImpl) Schedule(i AppointmentSchedulerInput) (AppointmentOutput, error) {
	p, err := u.professionalAcl.FindProfessionalByID(i.ProfessionalID)
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	s, err := u.serviceAcl.FindServiceByID(i.ServiceID)
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	c, err := u.findOrRegistrationCustomer(i)
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	id, err := id.Random()
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	a, err := NewAppointmentBuilder().
		WithAppointmentID(id).
		WithProfessionalID(p.ID).
		WithProfessionalName(p.Name).
		WithCustomerID(c.ID).
		WithCustomerName(c.Name).
		WithServiceID(s.ID).
		WithServiceName(s.Name).
		WithDate(i.Date).
		WithStartHour(i.StartHour).
		WithDuration(i.Duration).
		Build()

	if err != nil {
		return EmptyAppointmentOutput, err
	}

	appointments, err := u.repo.FindByDateAndStatus(Date(i.Date), StatusScheduled)
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	if !VerifyAvailability(a, appointments) {
		return EmptyAppointmentOutput, ErrBusyTime
	}

	u.repo.Save(a)

	e := event.New(
		EventAppointmentScheduled,
		event.WithHeader(event.HeaderAggregateID, a.ID.String()),
		event.WithPayload(i),
	)

	u.bus.Publish(e)

	return toAppointmentOutput(a), nil
}

func (u *appointmentScedulerImpl) findOrRegistrationCustomer(i AppointmentSchedulerInput) (Customer, error) {
	if len(i.CustomerID) == 0 && len(i.CustomerName) == 0 && len(i.CustomerPhone) == 0 {
		return Customer{}, ErrCustomerNotFound
	}

	if len(i.CustomerID) > 0 && (len(i.CustomerName) == 0 && len(i.CustomerPhone) == 0) {
		return u.customerAcl.FindCustomerByID(i.CustomerID)
	}

	return u.customerAcl.RequestCustomerRegistration(i.CustomerName, i.CustomerPhone)
}
