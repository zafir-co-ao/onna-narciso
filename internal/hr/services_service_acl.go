package hr

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
)

var ErrServiceNotFound = errors.New("service not found")

type ServicesServiceACL interface {
	FindServicesByIDs(ids []nanoid.ID) ([]Service, error)
}

type ServicesServiceACLFunc func(ids []nanoid.ID) ([]Service, error)

func (f ServicesServiceACLFunc) FindServicesByIDs(ids []nanoid.ID) ([]Service, error) {
	return f(ids)
}
