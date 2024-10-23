package scheduling

import "errors"

var ErrAppointmentNotFound = errors.New("appointment not found")

type AppointmentRepository interface {
	Save(appointment Appointment) error
	FindByID(id ID) (Appointment, error)
	FindByDate(date string) ([]Appointment, error)
	FindByDateAndStatus(date string, status Status) ([]Appointment, error)
	FindByWeekServiceAndProfessionals(week string, serviceID string, professionalsIDs []string) ([]Appointment, error)
}
