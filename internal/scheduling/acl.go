package scheduling

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
)

var (
	ErrInvalidService       = errors.New("invalid service")
	ErrServiceNotFound      = errors.New("service not found")
	ErrCustomerNotFound     = errors.New("customer not found")
	ErrProfessionalNotFound = errors.New("professional not found")
	ErrCustomerRegistration = errors.New("Error registering customer")
)

type ProfessionalsACL interface {
	FindProfessionalByID(id nanoid.ID) (Professional, error)
}

type CustomersACL interface {
	FindCustomerByID(id nanoid.ID) (Customer, error)
	RequestCustomerRegistration(name, phone string) (Customer, error)
}

type ProfessionalsACLFunc func(id nanoid.ID) (Professional, error)

func (f ProfessionalsACLFunc) FindProfessionalByID(id nanoid.ID) (Professional, error) {
	return f(id)
}

type ServicesService interface {
	FindServiceByID(id nanoid.ID) (Service, error)
}
