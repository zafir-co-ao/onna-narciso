package scheduling

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
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
	Hour           string
	Duration       int
}

type AppointmentScheduler interface {
	Schedule(i AppointmentSchedulerInput) (AppointmentOutput, error)
}

type appointmentScedulerImpl struct {
	repo            AppointmentRepository
	serviceACL      ServiceACL
	customerACL     CustomersACL
	professionalACL ProfessionalsACL
	bus             event.Bus
}

func NewAppointmentScheduler(
	repo AppointmentRepository,
	cacl CustomersACL,
	pacl ProfessionalsACL,
	sacl ServiceACL,
	bus event.Bus,
) AppointmentScheduler {
	return &appointmentScedulerImpl{
		repo:            repo,
		customerACL:     cacl,
		professionalACL: pacl,
		serviceACL:      sacl,
		bus:             bus,
	}
}

func (u *appointmentScedulerImpl) Schedule(i AppointmentSchedulerInput) (AppointmentOutput, error) {
	p, err := u.professionalACL.FindProfessionalByID(i.ProfessionalID)
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	s, err := u.serviceACL.FindServiceByID(i.ServiceID)
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	c, err := u.findOrRegistrationCustomer(i)
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	id := nanoid.New()
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
		WithHour(i.Hour).
		WithDuration(i.Duration).
		Build()

	if err != nil {
		return EmptyAppointmentOutput, err
	}

	appointments, err := u.repo.FindByDateStatusAndProfessional(date.Date(i.Date), StatusScheduled, p.ID)
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	if AppointmentsInterceptAny(a, appointments) {
		return EmptyAppointmentOutput, ErrBusyTime
	}

	err = u.repo.Save(a)
	if err != nil {
		return EmptyAppointmentOutput, err
	}

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
		return u.customerACL.FindCustomerByID(i.CustomerID)
	}

	return u.customerACL.RequestCustomerRegistration(i.CustomerName, i.CustomerPhone)
}
