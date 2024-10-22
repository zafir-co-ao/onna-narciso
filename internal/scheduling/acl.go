package scheduling

import "errors"

var (
	ErrServiceNotFound      = errors.New("service not found")
	ErrCustomerNotFound     = errors.New("customer not found")
	ErrProfessionalNotFound = errors.New("professional not found")
)

type ProfessionalAcl interface {
	FindProfessionalByID(id string) (Professional, error)
}

type ServiceAcl interface {
	FindServiceByID(id string) (Service, error)
}

type CustomerAcl interface {
	FindCustomerByID(id string) (Customer, error)
}

type ProfessionalAclFunc func(id string) (Professional, error)

func (f ProfessionalAclFunc) FindProfessionalByID(id string) (Professional, error) {
	return f(id)
}

type ServiceAclFunc func(id string) (Service, error)

func (f ServiceAclFunc) FindServiceByID(id string) (Service, error) {
	return f(id)
}

type CustomerAclFunc func(id string) (Customer, error)

func (f CustomerAclFunc) FindCustomerByID(id string) (Customer, error) {
	return f(id)
}