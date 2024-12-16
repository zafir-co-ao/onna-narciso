package scheduling

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

var (
	ErrCustomerNotFound     = errors.New("customer not found")
	ErrCustomerRegistration = errors.New("customer registration error")
)

type CRMServiceACL interface {
	FindCustomerByID(id nanoid.ID) (Customer, error)
	RequestCustomerRegistration(name, phone string) (Customer, error)
}

type internalCRMServiceACL struct {
	finder  crm.CustomerFinder
	creator crm.CustomerCreator
}

func NewInternalCRMServiceACL(finder crm.CustomerFinder, creator crm.CustomerCreator) CRMServiceACL {
	return &internalCRMServiceACL{
		finder:  finder,
		creator: creator,
	}
}

func (s *internalCRMServiceACL) FindCustomerByID(id nanoid.ID) (Customer, error) {
	c, err := s.finder.FindByID(id.String())
	if err != nil {
		return Customer{}, err
	}

	return Customer{
		ID:          nanoid.ID(c.ID),
		Name:        name.Name(c.Name),
		PhoneNumber: c.PhoneNumber,
	}, nil

}

func (s *internalCRMServiceACL) RequestCustomerRegistration(thename, phone string) (Customer, error) {

	i := crm.CustomerCreatorInput{
		Name:        thename,
		PhoneNumber: phone,
	}

	c, err := s.creator.Create(i)
	if err != nil {
		return Customer{}, err
	}

	return Customer{
		ID:          nanoid.ID(c.ID),
		Name:        name.Name(c.Name),
		PhoneNumber: c.PhoneNumber,
	}, nil
}
