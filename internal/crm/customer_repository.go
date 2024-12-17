package crm

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
)

var ErrCustomerNotFound = errors.New("customer not found")

type Repository interface {
	FindAll() ([]Customer, error)
	FindByID(id nanoid.ID) (Customer, error)
	Save(c Customer) error
}
