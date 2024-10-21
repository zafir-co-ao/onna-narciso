package scheduling

import (
	"time"
)

type AppointmentSchedulerDTO struct {
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
	Schedule(d AppointmentSchedulerDTO) (string, error)
}

type appointmentScedulerImpl struct {
	customerRepo     CustomerRepository
	appointmentRepo  AppointmentRepository
	professionalRepo ProfessionalRepository
	serviceRepo      ServiceRepository
}

func NewAppointmentScheduler(aRepo AppointmentRepository, cRepo CustomerRepository, pRepo ProfessionalRepository, sRepo ServiceRepository) AppointmentScheduler {
	return &appointmentScedulerImpl{
		appointmentRepo:  aRepo,
		customerRepo:     cRepo,
		professionalRepo: pRepo,
		serviceRepo:      sRepo,
	}
}

func (s *appointmentScedulerImpl) Schedule(d AppointmentSchedulerDTO) (string, error) {
	_, err := s.professionalRepo.Get(d.ProfessionalID)
	if err != nil {
		return "", err
	}

	_, err = s.serviceRepo.Get(d.ServiceID)
	if err != nil {
		return "", err
	}

	customer, err := s.getOrAddCustomer(d)
	if err != nil {
		return "", err
	}

	app, _ := NewAppointmentBuilder().
		WithAppointmentID("1").
		WithProfessionalID(d.ProfessionalID).
		WithCustomerID(customer.ID).
		WithServiceID(d.ServiceID).
		WithDate(d.Date).
		WithStartHour(d.StartHour).
		WithDuration(d.Duration).
		Build()

	appointments, _ := s.appointmentRepo.FindByDate(d.Date)
	if !VerifyAvailability(app, appointments) {
		return "", ErrBusyTime
	}

	s.appointmentRepo.Save(app)

	return "1", nil
}

func (s *appointmentScedulerImpl) getOrAddCustomer(d AppointmentSchedulerDTO) (Customer, error) {
	var customer Customer

	if len(d.CustomerID) > 0 {
		c, err := s.customerRepo.Get(d.CustomerID)
		return c, err
	}

	if len(d.CustomerName) > 0 && len(d.CustomerPhone) > 0 {
		customer = Customer{ID: "1000", Name: d.CustomerName, Phone: d.CustomerPhone}
		s.customerRepo.Save(customer)
	}

	return customer, nil
}

func VerifyAvailability(a Appointment, appointments []Appointment) bool {
	var isAvailable = true

	if len(appointments) == 0 {
		return isAvailable
	}

	for _, b := range appointments {
		if isNotAvailable(a, b) {
			isAvailable = false
			break
		}
	}

	return isAvailable

}

func isNotAvailable(a, b Appointment) bool {
	startTimeA, _ := time.Parse("15:04", a.Start)
	endTimeA, _ := time.Parse("15:04", a.End)
	startTimeB, _ := time.Parse("15:04", b.Start)
	endTimeB, _ := time.Parse("15:04", b.End)

	if startTimeA.Equal(startTimeB) {
		return true
	}

	if startTimeA.Before(startTimeB) && endTimeA.After(startTimeB) {
		return true
	}

	if startTimeA.After(startTimeB) && startTimeA.Before(endTimeB) {
		return true
	}

	return false
}
