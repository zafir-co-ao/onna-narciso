package scheduling

import "errors"

var ErrCustomerNotFound = errors.New("Client not found")

type CustomerRepository interface {
	Get(id string) (Customer, error)
	Save(c Customer) error
}
