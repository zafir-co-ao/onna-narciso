package scheduling

import "errors"

var (
	ErrServiceNotFound      = errors.New("service not found")
	ErrCustomerNotFound     = errors.New("customer not found")
	ErrProfessionalNotFound = errors.New("professional not found")
	ErrCustomerRegistration = errors.New("Error registering customer")
)

type ProfessionalACL interface {
	FindProfessionalByID(id string) (Professional, error)
}

type ServiceACL interface {
	FindServiceByID(id string) (Service, error)
}

type CustomerACL interface {
	FindCustomerByID(id string) (Customer, error)
	RequestCustomerRegistration(name, phone string) (Customer, error)
}

type ProfessionalACLFunc func(id string) (Professional, error)

func (f ProfessionalACLFunc) FindProfessionalByID(id string) (Professional, error) {
	return f(id)
}

type ServiceACLFunc func(id string) (Service, error)

func (f ServiceACLFunc) FindServiceByID(id string) (Service, error) {
	return f(id)
}
