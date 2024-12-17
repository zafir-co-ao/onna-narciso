package scheduling

import (
	"errors"
	"fmt"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/duration"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

var (
	ErrInvalidService  = errors.New("invalid service")
	ErrServiceNotFound = errors.New("service not found")
)

type ServicesServiceACL interface {
	FindServiceByID(id nanoid.ID) (Service, error)
}

type ServicesServiceACLFunc func(id nanoid.ID) (Service, error)

func (f ServicesServiceACLFunc) FindServiceByID(id nanoid.ID) (Service, error) {
	return f(id)
}

func NewInternalServicesACL(finder services.ServiceFinder) ServicesServiceACL {
	return &internalservicesServiceACL{finder}
}

type internalservicesServiceACL struct {
	finder services.ServiceFinder
}

func (i *internalservicesServiceACL) FindServiceByID(id nanoid.ID) (Service, error) {
	s, err := i.finder.FindByID(id.String())

	if err != nil {
		return Service{}, fmt.Errorf("%w: %w", ErrServiceNotFound, err)
	}

	return Service{
		ID:       nanoid.ID(s.ID),
		Name:     name.Name(s.Name),
		Duration: duration.Duration(s.Duration),
	}, nil
}
