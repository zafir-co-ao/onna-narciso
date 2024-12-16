package integration

import (
	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
)

type CRMServiceACL interface {
	GetCustomer(id string) (crm.CustomerOutput, error)
	GetBirthdayCustomers(d date.Date) ([]crm.CustomerOutput, error)
}

type internalCRMServiceACL struct {
	finder crm.CustomerFinder
}

func NewInternalCRMServiceACL(finder crm.CustomerFinder) CRMServiceACL {
	return &internalCRMServiceACL{finder: finder}
}

func (s *internalCRMServiceACL) GetBirthdayCustomers(d date.Date) ([]crm.CustomerOutput, error) {

	day := d.String()[4:]

	customers, err := s.finder.FindAll()
	if err != nil {
		return nil, err
	}

	filtered := xslices.Filter(customers, func(c crm.CustomerOutput) bool {
		return c.BirthDate[4:] == day
	})

	return filtered, nil
}

func (s *internalCRMServiceACL) GetCustomer(id string) (crm.CustomerOutput, error) {
	return s.finder.FindByID(id)
}
