package crm

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
)

type inmemCustomerRepository struct {
	shared.BaseRepository[Customer]
}

func NewInmemRepository(c ...Customer) Repository {
	return &inmemCustomerRepository{
		BaseRepository: shared.NewBaseRepository[Customer](c...),
	}
}

func (c *inmemCustomerRepository) FindAll() ([]Customer, error) {
	var customers []Customer

	for _, customer := range c.Data {
		customers = append(customers, customer)
	}

	return customers, nil
}

func (c *inmemCustomerRepository) FindByID(id nanoid.ID) (Customer, error) {
	if _, ok := c.Data[id]; !ok {
		return Customer{}, ErrCustomerNotFound
	}

	return c.Data[id], nil
}

func (c *inmemCustomerRepository) FindByNif(nif Nif) (Customer, error) {
	for _, customer := range c.Data {
		if customer.Nif == nif {
			return customer, ErrNifAlreadyUsed
		}
	}

	return Customer{}, nil
}

func (c *inmemCustomerRepository) FindByEmail(e Email) (Customer, error) {
	for _, customer := range c.Data {
		if customer.Email == e {
			return customer, nil
		}
	}

	return Customer{}, ErrCustomerNotFound
}

func(c *inmemCustomerRepository) FindByPhoneNumber(p PhoneNumber) (Customer, error) {
	for _, customer := range c.Data {
		if customer.PhoneNumber == p {
			return customer, nil
		}
	}

	return Customer{}, ErrCustomerNotFound
}

func (c *inmemCustomerRepository) Save(customer Customer) error {
	c.Data[customer.ID] = customer
	return nil
}
