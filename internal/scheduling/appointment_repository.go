package scheduling

import "errors"

var ErrAppointmentNotFound = errors.New("appointment not found")

type AppointmentRepository interface {
	Get(id string) (Appointment, error)
	Save(appointment Appointment) error
	FindByDate(date string) ([]Appointment, error)
	FindByWeekServiceAndProfessionals(week string, serviceID string, professionalsIDs []string) ([]Appointment, error)
}
