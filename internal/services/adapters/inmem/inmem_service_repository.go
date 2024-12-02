package inmem

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
)

type inmemServiceRepositoryImpl struct {
	shared.BaseRepository[services.Service]
}

func NewServiceRepository(s ...services.Service) services.Repository {
	return &inmemServiceRepositoryImpl{
		BaseRepository: shared.NewBaseRepository[services.Service](s...),
	}
}

func (s *inmemServiceRepositoryImpl) FindByID(id nanoid.ID) (services.Service, error) {
	return s.Data[id], nil
}

func (s *inmemServiceRepositoryImpl) Save(service services.Service) error {
	s.Data[service.ID] = service
	return nil
}
