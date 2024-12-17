package sessions

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
)

var ErrServiceNotFound = errors.New("service not found")

var EmptyServices = make([]SessionService, 0)

type ServicesServiceACL interface {
	FindByIDs(i []nanoid.ID) ([]SessionService, error)
}

type ServicesACLFunc func(i []nanoid.ID) ([]SessionService, error)

func (f ServicesACLFunc) FindByIDs(i []nanoid.ID) ([]SessionService, error) {
	return f(i)
}

type internalServicesServiceACL struct {
	finder services.ServiceFinder
}

func NewServicesServiceACL(finder services.ServiceFinder) ServicesServiceACL {
	return &internalServicesServiceACL{finder}
}

func (i *internalServicesServiceACL) FindByIDs(ids []nanoid.ID) ([]SessionService, error) {
	sids := xslices.Map(ids, func(id nanoid.ID) string { return id.String() })

	s, err := i.finder.FindByIDs(sids)
	if err != nil {
		return EmptyServices, ErrServiceNotFound
	}

	return xslices.Map(s, func(s services.ServiceOutput) SessionService {
		return SessionService{
			ServiceID:   nanoid.ID(s.ID),
			ServiceName: s.Name,
		}
	}), nil
}
