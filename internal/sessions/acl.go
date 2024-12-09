package sessions

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
)

var (
	ErrServiceNotFound     = errors.New("service not found")
	ErrAppointmentNotFound = errors.New("appointment not found")
	ErrAppointmentCanceled = errors.New("appointment canceled")
)

var EmptyServices = make([]SessionService, 0, 0)

type ServicesACL interface {
	FindByIDs(i []nanoid.ID) ([]SessionService, error)
}

type ServicesACLFunc func(i []nanoid.ID) ([]SessionService, error)

func (f ServicesACLFunc) FindByIDs(i []nanoid.ID) ([]SessionService, error) {
	return f(i)
}

type AppointmentsACL interface {
	FindByID(i nanoid.ID) (Appointment, error)
}

type AppointmentsACLFunc func(i nanoid.ID) (Appointment, error)

func (f AppointmentsACLFunc) FindByID(i nanoid.ID) (Appointment, error) {
	return f(i)
}
