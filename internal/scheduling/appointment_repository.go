package scheduling

import "errors"

var ErrAppointmentNotFound = errors.New("appointment not found")

type AppointmentRepository interface {
	Save(appointment Appointment) error
	FindByID(id string) (Appointment, error)
	FindByDate(date string) ([]Appointment, error)
	FindByWeekServiceAndProfessionals(week string, serviceID string, professionalsIDs []string) ([]Appointment, error)
}
