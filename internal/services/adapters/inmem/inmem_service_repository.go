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

func (s *inmemServiceRepositoryImpl) FindAll() ([]services.Service, error) {
	var services []services.Service

	for _, s := range s.Data {
		services = append(services, s)
	}

	return services, nil
}

func (s *inmemServiceRepositoryImpl) FindByID(id nanoid.ID) (services.Service, error) {
	if _, ok := s.Data[id]; !ok {
		return services.Service{}, services.ErrServiceNotFound
	}

	return s.Data[id], nil
}

func (s *inmemServiceRepositoryImpl) Save(service services.Service) error {
	s.Data[service.ID] = service
	return nil
}
