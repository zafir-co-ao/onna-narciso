package services

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
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
	_ids := xslices.Map(ids, shared.StringToNanoid)

	s, err := u.repo.FindByIDs(_ids)

	if err != nil {
		return []ServiceOutput{}, err
	}

	return xslices.Map(s, toServiceOutput), nil
}
