package services

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
)

type inmemServiceRepositoryImpl struct {
	shared.BaseRepository[Service]
}

func NewInmemRepository(s ...Service) Repository {
	return &inmemServiceRepositoryImpl{
		BaseRepository: shared.NewBaseRepository[Service](s...),
	}
}

func (s *inmemServiceRepositoryImpl) FindAll() ([]Service, error) {
	var services []Service

	for _, s := range s.Data {
		services = append(services, s)
	}

	return services, nil
}

func (s *inmemServiceRepositoryImpl) FindByID(id nanoid.ID) (Service, error) {
	return s.Data[id], nil
}

func (s *inmemServiceRepositoryImpl) Save(service Service) error {
	s.Data[service.ID] = service
	return nil
}
