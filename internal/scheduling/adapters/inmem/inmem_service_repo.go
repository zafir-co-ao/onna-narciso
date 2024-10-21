package inmem

import "github.com/zafir-co-ao/onna-narciso/internal/scheduling"

type serviceRepo struct {
	data map[string]scheduling.Service
}

func NewServiceRepository() scheduling.ServiceRepository {
	return &serviceRepo{
		data: make(map[string]scheduling.Service),
	}
}

func (r *serviceRepo) Get(id string) (scheduling.Service, error) {
	if service, ok := r.data[id]; ok {
		return service, nil
	}
	return scheduling.Service{}, scheduling.ErrServiceNotFound
}

func (r *serviceRepo) Save(s scheduling.Service) error {
	r.data[s.ID] = s
	return nil
}
