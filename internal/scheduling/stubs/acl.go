package stubs

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
)

func NewProfessionalsACL() scheduling.ProfessionalsACL {
	f := func(id nanoid.ID) (scheduling.Professional, error) {
		for _, p := range testdata.Professionals {
			if p.ID == id {
				return p, nil
			}
		}

		return scheduling.Professional{}, scheduling.ErrProfessionalNotFound
	}

	return scheduling.ProfessionalsACLFunc(f)
}

func NewServicesACL() scheduling.ServicesService {
	f := func(id nanoid.ID) (scheduling.Service, error) {
		for _, s := range testdata.Services {
			if s.ID == id {
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

func (c customerACLStub) FindCustomerByID(id nanoid.ID) (scheduling.Customer, error) {
	for _, c := range testdata.Customers {
		if c.ID == id {
			return c, nil
		}
	}

	return scheduling.Customer{}, scheduling.ErrCustomerNotFound
}

func (c customerACLStub) RequestCustomerRegistration(name string, phone string) (scheduling.Customer, error) {
	if name == "" || phone == "" {
		return scheduling.Customer{}, scheduling.ErrCustomerRegistration
	}

	return scheduling.Customer{ID: "1", Name: scheduling.Name(name), PhoneNumber: phone}, nil
}
