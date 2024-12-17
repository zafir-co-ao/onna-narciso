package crm

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
)

type inmemCustomerRepositoryImpl struct {
	shared.BaseRepository[Customer]
}

func NewInmemRepository(c ...Customer) Repository {
	return &inmemCustomerRepositoryImpl{BaseRepository: shared.NewBaseRepository[Customer](c...)}
}

func (r *inmemCustomerRepositoryImpl) FindAll() ([]Customer, error) {
	var customers []Customer
	for _, c := range r.Data {
		customers = append(customers, c)
	}

	return customers, nil
}

func (r *inmemCustomerRepositoryImpl) FindByID(id nanoid.ID) (Customer, error) {
	if _, ok := r.Data[id]; !ok {
		return Customer{}, ErrCustomerNotFound
	}

	return r.Data[id], nil
}

func (r *inmemCustomerRepositoryImpl) Save(c Customer) error {
	r.Data[c.ID] = c
	return nil
}
