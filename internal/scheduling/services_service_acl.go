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

func NewInternalServicesACL(r services.Repository) ServicesServiceACL {
	return &internalservicesServiceACL{r}
}

type internalservicesServiceACL struct {
	repository services.Repository
}

func (i *internalservicesServiceACL) FindServiceByID(id nanoid.ID) (Service, error) {

	s, err := i.repository.FindByID(id)

	if err != nil {
		return Service{}, fmt.Errorf("%w: %w", ErrServiceNotFound, err)
	}

	o := Service{
		ID:       s.ID,
		Name:     name.Name(s.Name),
		Duration: duration.Duration(s.Duration.Value()),
	}

	return o, nil
}
