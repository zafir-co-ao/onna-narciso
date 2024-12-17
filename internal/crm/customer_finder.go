package crm

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
)

type CustomerFinder interface {
	FindAll() ([]CustomerOutput, error)
	FindByID(id string) (CustomerOutput, error)
}

type finderImpl struct {
	repo Repository
}

func NewCustomerFinder(repo Repository) CustomerFinder {
	return &finderImpl{repo}
}

func (u *finderImpl) FindAll() ([]CustomerOutput, error) {

	c, err := u.repo.FindAll()
	if err != nil {
		return []CustomerOutput{}, err
	}

	return xslices.Map(c, toCustomerOutput), nil
}

func (u *finderImpl) FindByID(id string) (CustomerOutput, error) {
	c, err := u.repo.FindByID(nanoid.ID(id))

	if err != nil {
		return CustomerOutput{}, err
	}

	return toCustomerOutput(c), nil
}
