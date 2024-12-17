package stubs

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	_name "github.com/zafir-co-ao/onna-narciso/internal/shared/name"
	testdata "github.com/zafir-co-ao/onna-narciso/test_data"
)

type crmServiceACLStub struct{}

func NewCRMServiceACL() scheduling.CRMServiceACL {
	return crmServiceACLStub{}
}

func (c crmServiceACLStub) FindCustomerByID(id nanoid.ID) (scheduling.Customer, error) {
	for _, c := range testdata.Customers {
		if c.ID == id {
			return c, nil
		}
	}

	return scheduling.Customer{}, scheduling.ErrCustomerNotFound
}

func (c crmServiceACLStub) RequestCustomerRegistration(name string, phone string) (scheduling.Customer, error) {
	if name == "" || phone == "" {
		return scheduling.Customer{}, scheduling.ErrCustomerRegistration
	}

	return scheduling.Customer{ID: "1", Name: _name.Name(name), PhoneNumber: phone}, nil
}
