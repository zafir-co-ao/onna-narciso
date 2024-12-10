package crm

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
)

var (
	ErrCustomerNotFound       = errors.New("customer not found")
	ErrNifAlreadyUsed         = errors.New("nif already used")
	ErrEmailAlreadyUsed       = errors.New("email already used")
	ErrPhoneNumberAlreadyUsed = errors.New("phone number already used")
)

type Repository interface {
	FindAll() ([]Customer, error)
	FindByID(id nanoid.ID) (Customer, error)
	FindByNif(n Nif) (Customer, error)
	FindByEmail(e Email) (Customer, error)
	FindByPhoneNumber(p PhoneNumber) (Customer, error)
	Save(c Customer) error
}
