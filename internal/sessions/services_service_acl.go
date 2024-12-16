package sessions

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
)

var (
	ErrServiceNotFound     = errors.New("service not found")
	ErrAppointmentNotFound = errors.New("appointment not found")
	ErrAppointmentCanceled = errors.New("appointment canceled")
	ErrAppointmentClosed   = errors.New("appointment closed")
)

var EmptyServices = make([]SessionService, 0)

type ServicesServiceACL interface {
	FindByIDs(i []nanoid.ID) ([]SessionService, error)
}

type ServicesACLFunc func(i []nanoid.ID) ([]SessionService, error)

func (f ServicesACLFunc) FindByIDs(i []nanoid.ID) ([]SessionService, error) {
	return f(i)
}
