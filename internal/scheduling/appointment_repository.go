package scheduling

import "errors"

var ErrAppointmentNotFound = errors.New("appointment not found")

type AppointmentRepository interface {
	Save(a Appointment) error
	FindByID(id ID) (Appointment, error)
	FindByDate(date Date) ([]Appointment, error)
	FindByDateAndStatus(date Date, status Status) ([]Appointment, error)
	FindByWeekServiceAndProfessionals(week string, serviceID string, professionalsIDs []string) ([]Appointment, error)
}
