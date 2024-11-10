package scheduling

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
)

var ErrAppointmentNotFound = errors.New("appointment not found")

type AppointmentRepository interface {
	Save(a Appointment) error
	FindByID(id nanoid.ID) (Appointment, error)
	FindByDate(date Date) ([]Appointment, error)
	FindByDateStatusAndProfessional(date Date, status Status, id nanoid.ID) ([]Appointment, error)
	FindByWeekServiceAndProfessionals(date Date, serviceID nanoid.ID, professionalsIDs []nanoid.ID) ([]Appointment, error)
}
