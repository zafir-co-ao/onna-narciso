package scheduling

import (
	"fmt"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/duration"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

func NewMonoServicesService(r services.Repository) ServicesService {
	return &servicesServiceImpl{r}
}

type servicesServiceImpl struct {
	repository services.Repository
}

func (i *servicesServiceImpl) FindServiceByID(id nanoid.ID) (Service, error) {

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
