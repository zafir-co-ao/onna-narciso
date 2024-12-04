package services

import "github.com/kindalus/godx/pkg/xslices"

type ServiceFinder interface {
	Find() ([]ServiceOutput, error)
}

type serviceFinderImpl struct {
	repo Repository
}

func NewServiceFinder(repo Repository) ServiceFinder {
	return &serviceFinderImpl{repo}
}

func (u *serviceFinderImpl) Find() ([]ServiceOutput, error) {
	s, err := u.repo.FindAll()

	if err != nil {
		return []ServiceOutput{}, err
	}

	return xslices.Map(s, toServiceOutput), nil
}
