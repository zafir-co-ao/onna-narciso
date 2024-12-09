package crm

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/nif"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
)

type inmemCustomerRepository struct {
	shared.BaseRepository[Customer]
}

func NewInmemRepository(c ...Customer) Repository {
	return &inmemCustomerRepository{
		BaseRepository: shared.NewBaseRepository[Customer](c ...),
	}
}

func (c *inmemCustomerRepository) FindByID(id nanoid.ID) (Customer, error) {
	if _, ok := c.Data[id]; !ok {
		return Customer{}, ErrCustomerNotFound
	}

	return c.Data[id], nil
}

func (c *inmemCustomerRepository) FindByNif(nif nif.Nif) (Customer, error) {
	for _, customer := range c.Data {
		if customer.Nif == nif {
			return customer, ErrNifAlreadyUsed
		}
	}

	return Customer{}, nil
}

func (c *inmemCustomerRepository) Save(customer Customer) error {
	c.Data[customer.ID] = customer
	return nil
}
