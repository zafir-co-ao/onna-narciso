package scheduling

import "errors"

var (
	ErrServiceNotFound      = errors.New("service not found")
	ErrCustomerNotFound     = errors.New("customer not found")
	ErrProfessionalNotFound = errors.New("professional not found")
	ErrCustomerRegistration = errors.New("Error registering customer")
)

type ProfessionalsACL interface {
	FindProfessionalByID(id string) (Professional, error)
}

type ServiceACL interface {
	FindServiceByID(id string) (Service, error)
}

type CustomersACL interface {
	FindCustomerByID(id string) (Customer, error)
	RequestCustomerRegistration(name, phone string) (Customer, error)
}

type ProfessionalsACLFunc func(id string) (Professional, error)

func (f ProfessionalsACLFunc) FindProfessionalByID(id string) (Professional, error) {
	return f(id)

}

type ServicesACLFunc func(id string) (Service, error)

func (f ServicesACLFunc) FindServiceByID(id string) (Service, error) {
	return f(id)
}
