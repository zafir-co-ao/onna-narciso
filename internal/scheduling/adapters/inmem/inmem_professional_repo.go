package inmem

import "github.com/zafir-co-ao/onna-narciso/internal/scheduling"

type professionalRepo struct {
	data map[string]scheduling.Professional
}

func NewProfessionalRepository() scheduling.ProfessionalRepository {
	return &professionalRepo{
		data: make(map[string]scheduling.Professional),
	}
}

func (r *professionalRepo) Get(id string) (scheduling.Professional, error) {
	if p, ok := r.data[id]; ok {
		return p, nil
	}
	return scheduling.EmptyProfessional, scheduling.ErrProfessionalNotFound
}

func (r *professionalRepo) Save(p scheduling.Professional) error {
	r.data[p.ID] = p
	return nil
}
