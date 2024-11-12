package session

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
)

var ErrServiceNotFound = errors.New("service not found")

var EmptyServices = []Service{}

type ServiceACL interface {
	FindByIDs(i []nanoid.ID) ([]Service, error)
}

type ServiceACLFunc func(i []nanoid.ID) ([]Service, error)

func (f ServiceACLFunc) FindByIDs(i []nanoid.ID) ([]Service, error) {
	return f(i)
}
