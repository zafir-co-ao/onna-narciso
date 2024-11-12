package stubs

import (
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
)

func NewProfessionalsACL() scheduling.ProfessionalsACL {
	f := func(id string) (scheduling.Professional, error) {
		for _, p := range testdata.Professionals {
			if p.ID.String() == id {
				return p, nil
			}
		}

		return scheduling.Professional{}, scheduling.ErrProfessionalNotFound
	}

	return scheduling.ProfessionalsACLFunc(f)
}

func NewServicesACL() scheduling.ServiceACL {
	f := func(id string) (scheduling.Service, error) {
		for _, s := range testdata.Services {
			if s.ID.String() == id {
				return s, nil
			}
		}

		return scheduling.Service{}, scheduling.ErrServiceNotFound
	}

	return scheduling.ServicesACLFunc(f)
}

func NewCustomersACL() scheduling.CustomersACL {
	return customerACLStub{}
}

type customerACLStub struct{}

func (c customerACLStub) FindCustomerByID(id string) (scheduling.Customer, error) {
	switch id {
	case "1":
		return scheduling.Customer{ID: "1", Name: "João Silva"}, nil
	case "2":
		return scheduling.Customer{ID: "2", Name: "Maria Oliveira"}, nil
	case "3":
		return scheduling.Customer{ID: "3", Name: "Carlos Ferreira"}, nil
	default:
		return scheduling.Customer{}, scheduling.ErrCustomerNotFound
	}
}

func (c customerACLStub) RequestCustomerRegistration(name string, phone string) (scheduling.Customer, error) {
	if name == "" || phone == "" {
		return scheduling.Customer{}, scheduling.ErrCustomerRegistration
	}

	return scheduling.Customer{ID: "1", Name: scheduling.Name(name), PhoneNumber: phone}, nil
}
