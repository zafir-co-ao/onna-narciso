package crm

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
)

type CustomerFinder interface {
	FindAll() ([]CustomerOutput, error)
	FindByID(id string) (CustomerOutput, error)
}

type customerFinderImpl struct {
	repo Repository
}

func NewCustomerFinder(repo Repository) CustomerFinder {
	return &customerFinderImpl{repo}
}

func (u *customerFinderImpl) FindAll() ([]CustomerOutput, error) {

	c, err := u.repo.FindAll()
	if err != nil {
		return []CustomerOutput{}, err
	}

	return xslices.Map(c, toCustomerOutput), nil
}

func (u *customerFinderImpl) FindByID(id string) (CustomerOutput, error) {
	c, err := u.repo.FindByID(nanoid.ID(id))

	if err != nil {
		return CustomerOutput{}, err
	}

	return toCustomerOutput(c), nil
}
