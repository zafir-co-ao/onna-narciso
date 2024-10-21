package inmem

import "github.com/zafir-co-ao/onna-narciso/internal/scheduling"

type customerRepo struct {
	data map[string]scheduling.Customer
}

func NewCustomerRepository() scheduling.CustomerRepository {
	return &customerRepo{
		data: make(map[string]scheduling.Customer),
	}
}

func (r *customerRepo) Get(id string) (scheduling.Customer, error) {
	customer, ok := r.data[id]
	if !ok {
		return scheduling.EmptyCustomer, scheduling.ErrCustomerNotFound
	}
	return customer, nil
}

func (r *customerRepo) Save(c scheduling.Customer) error {
	r.data[c.ID] = c
	return nil
}
