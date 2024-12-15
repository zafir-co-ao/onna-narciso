package scheduling

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
)

var ErrAppointmentNotFound = errors.New("appointment not found")

type AppointmentRepository interface {
	Save(a Appointment) error
	FindByID(id nanoid.ID) (Appointment, error)
	FindByDate(date date.Date) ([]Appointment, error)
	FindActivesByDateAndProfessional(date date.Date, professionalID nanoid.ID) ([]Appointment, error)
	FindByWeekServiceAndProfessionals(date date.Date, serviceID nanoid.ID, professionalsIDs []nanoid.ID) ([]Appointment, error)
}
