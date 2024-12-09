package crm

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
)

var (
	ErrCustomerNotFound = errors.New("customer not found")
	ErrNifAlreadyUsed   = errors.New("nif already used")
)

type Repository interface {
	FindAll() ([]Customer, error)
	FindByID(id nanoid.ID) (Customer, error)
	FindByNif(nif Nif) (Customer, error)
	Save(c Customer) error
}
