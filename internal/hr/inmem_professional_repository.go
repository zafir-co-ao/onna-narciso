package hr

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
)

type inmemProfessionalRepositoryImpl struct {
	shared.BaseRepository[Professional]
}

func NewInmemProfessionalRepository(p ...Professional) Repository {
	return &inmemProfessionalRepositoryImpl{BaseRepository: shared.NewBaseRepository[Professional](p...)}
}

func (r *inmemProfessionalRepositoryImpl) FindByID(id nanoid.ID) (Professional, error) {
	if _, ok := r.Data[id]; !ok {
		return Professional{}, ErrProfessionalNotFound
	}

	return r.Data[id], nil
}

func (r *inmemProfessionalRepositoryImpl) FindAll() ([]Professional, error) {
	var professionals []Professional
	for _, p := range r.Data {
		professionals = append(professionals, p)
	}

	return professionals, nil
}

func (r *inmemProfessionalRepositoryImpl) Save(p Professional) error {
	r.Data[p.ID] = p
	return nil
}
