package services

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
)

type inmemServiceRepositoryImpl struct {
	shared.BaseRepository[Service]
}

func NewInmemRepository(s ...Service) Repository {
	return &inmemServiceRepositoryImpl{BaseRepository: shared.NewBaseRepository[Service](s...)}
}

func (r *inmemServiceRepositoryImpl) FindAll() ([]Service, error) {
	var services []Service

	for _, s := range r.Data {
		services = append(services, s)
	}

	return services, nil
}

func (r *inmemServiceRepositoryImpl) FindByID(id nanoid.ID) (Service, error) {
	if _, ok := r.Data[id]; !ok {
		return Service{}, ErrServiceNotFound
	}

	return r.Data[id], nil
}

func (r *inmemServiceRepositoryImpl) Save(service Service) error {
	r.Data[service.ID] = service
	return nil
}
