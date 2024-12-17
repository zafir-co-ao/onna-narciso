package services

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
)

type ServiceFinder interface {
	FindAll() ([]ServiceOutput, error)
	FindByID(id string) (ServiceOutput, error)
	FindByIDs(ids []string) ([]ServiceOutput, error)
}

type serviceFinderImpl struct {
	repo Repository
}

func NewServiceFinder(repo Repository) ServiceFinder {
	return &serviceFinderImpl{repo}
}

func (u *serviceFinderImpl) FindAll() ([]ServiceOutput, error) {
	s, err := u.repo.FindAll()

	if err != nil {
		return []ServiceOutput{}, err
	}

	return xslices.Map(s, toServiceOutput), nil
}

func (u *serviceFinderImpl) FindByID(id string) (ServiceOutput, error) {
	s, err := u.repo.FindByID(nanoid.ID(id))

	if err != nil {
		return ServiceOutput{}, err
	}

	return toServiceOutput(s), err
}

func (u *serviceFinderImpl) FindByIDs(ids []string) ([]ServiceOutput, error) {
	nids := xslices.Map(ids, func(id string) nanoid.ID { return nanoid.ID(id) })

	s, err := u.repo.FindByIDs(nids)

	if err != nil {
		return []ServiceOutput{}, err
	}

	return xslices.Map(s, toServiceOutput), nil
}
