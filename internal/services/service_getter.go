package services

import "github.com/kindalus/godx/pkg/nanoid"

type ServiceGetter interface {
	Get(id string) (ServiceOutput, error)
}

type serviceGetterImpl struct {
	repo Repository
}

func NewServiceGetter(repo Repository) ServiceGetter {
	return &serviceGetterImpl{repo}
}

func (u *serviceGetterImpl) Get(id string) (ServiceOutput, error) {
	s, err := u.repo.FindByID(nanoid.ID(id))

	if err != nil {
		return ServiceOutput{}, err
	}

	return toServiceOutput(s), err
}
