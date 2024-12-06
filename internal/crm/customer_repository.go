package crm

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/nif"
)

var (
	ErrCustomerNotFound = errors.New("customer not found")
	ErrNifAlreadyUsed   = errors.New("nif already used")
)

type Repository interface {
	FindByID(id nanoid.ID) (Customer, error)
	FindByNif(nif nif.Nif) (Customer, error)
	Save(c Customer) error
}
