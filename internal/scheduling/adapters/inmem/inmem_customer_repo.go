package inmem

import "github.com/zafir-co-ao/onna-narciso/internal/scheduling"

type customerRepo struct {
	customers map[string]scheduling.Customer
}

func NewCustomerRepository() scheduling.CustomerRepository {
	return &customerRepo{
		customers: make(map[string]scheduling.Customer),
	}
}

func (r *customerRepo) Get(id string) (scheduling.Customer, error) {
	customer, ok := r.customers[id]
	if !ok {
		return scheduling.Customer{}, scheduling.ErrCustomerNotFound
	}
	return customer, nil
}

func (r *customerRepo) Save(c scheduling.Customer) error {
	r.customers[c.ID] = c
	return nil
}
