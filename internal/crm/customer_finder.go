package crm

import "github.com/kindalus/godx/pkg/xslices"

type CustomerFinder interface {
	Find() ([]CustomerOutput, error)
}

type customerFinderImpl struct {
	repo Repository
}

func NewCustomerFinder(repo Repository) CustomerFinder {
	return &customerFinderImpl{repo}
}

func (u *customerFinderImpl) Find() ([]CustomerOutput, error) {

	c, err := u.repo.FindAll()
	if err != nil {
		return []CustomerOutput{}, err
	}

	return xslices.Map(c, toCustomerOutput), nil
}
