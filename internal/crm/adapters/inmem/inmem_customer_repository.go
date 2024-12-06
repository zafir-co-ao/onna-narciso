package inmem

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/nif"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
)

type inmemCustomerRepository struct {
	shared.BaseRepository[crm.Customer]
}

func NewCustomerRepository(c ...crm.Customer) crm.Repository {
	return &inmemCustomerRepository{
		BaseRepository: shared.NewBaseRepository[crm.Customer](c ...),
	}
}

func (c *inmemCustomerRepository) FindByID(id nanoid.ID) (crm.Customer, error) {
	if _, ok := c.Data[id]; !ok {
		return crm.Customer{}, crm.ErrCustomerNotFound
	}

	return c.Data[id], nil
}

func (c *inmemCustomerRepository) FindByNif(nif nif.Nif) (crm.Customer, error) {
	for _, customer := range c.Data {
		if customer.Nif == nif {
			return customer, crm.ErrNifAlreadyUsed
		}
	}

	return crm.Customer{}, nil
}

func (c *inmemCustomerRepository) Save(customer crm.Customer) error {
	c.Data[customer.ID] = customer
	return nil
}
