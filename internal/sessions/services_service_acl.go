package sessions

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
)

var ErrServiceNotFound = errors.New("service not found")

var EmptyServices = make([]SessionService, 0)

type ServicesServiceACL interface {
	FindByIDs(id []nanoid.ID) ([]SessionService, error)
}

type ServicesACLFunc func(id []nanoid.ID) ([]SessionService, error)

func (f ServicesACLFunc) FindByIDs(id []nanoid.ID) ([]SessionService, error) {
	return f(id)
}

type internalServicesServiceACL struct {
	finder services.ServiceFinder
}

func NewServicesServiceACL(finder services.ServiceFinder) ServicesServiceACL {
	return &internalServicesServiceACL{finder}
}

func (i *internalServicesServiceACL) FindByIDs(ids []nanoid.ID) ([]SessionService, error) {
	_ids := xslices.Map(ids, shared.NanoidToString)

	s, err := i.finder.FindByIDs(_ids)
	if err != nil {
		return EmptyServices, ErrServiceNotFound
	}

	return xslices.Map(s, func(s services.ServiceOutput) SessionService {
		return SessionService{
			ID:       nanoid.ID(s.ID),
			Name:     s.Name,
			Price:    s.Price,
			Discount: s.Discount,
		}
	}), nil
}
