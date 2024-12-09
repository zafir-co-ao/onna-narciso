package crm

import "github.com/kindalus/godx/pkg/nanoid"

type CustomerGetter interface {
	Get(id string) (CustomerOutput, error)
}

type customerGetterImpl struct {
	repo Repository
}

func NewCustomerGetter(repo Repository) CustomerGetter {
	return &customerGetterImpl{repo}
}

func (u *customerGetterImpl) Get(id string) (CustomerOutput, error) {
	c, err := u.repo.FindByID(nanoid.ID(id))

	if err != nil {
		return CustomerOutput{}, err
	}

	return toCustomerOutput(c), nil
}
