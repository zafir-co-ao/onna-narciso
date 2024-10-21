package scheduling

import (
	"time"
)

type AppointmentSchedulerDTO struct {
	ID             string
	ProfessionalID string
	CustomerID     string
	ServiceID      string
	Date           string
	StartHour      string
	Duration       int
}

type AppointmentScheduler interface {
	Schedule(d AppointmentSchedulerDTO) (string, error)
}

type appointmentScedulerImpl struct {
	repo AppointmentRepository
}

func NewAppointmentScheduler(repo AppointmentRepository) AppointmentScheduler {
	return &appointmentScedulerImpl{
		repo: repo,
	}
}

func (s *appointmentScedulerImpl) Schedule(d AppointmentSchedulerDTO) (string, error) {
	appointments, _ := s.repo.FindByDate(d.Date)

	app, _ := NewAppointmentBuilder().
		WithAppointmentID("1").
		WithProfessionalID(d.ProfessionalID).
		WithCustomerID(d.CustomerID).
		WithServiceID(d.ServiceID).
		WithDate(d.Date).
		WithStartHour(d.StartHour).
		WithDuration(d.Duration).
		Build()

	if !VerifyAvailability(app, appointments) {
		return "", ErrBusyTime
	}

	s.repo.Save(app)

	return "1", nil
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
