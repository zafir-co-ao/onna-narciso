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

type finderImpl struct {
	repo Repository
}

func NewServiceFinder(repo Repository) ServiceFinder {
	return &finderImpl{repo}
}

func (u *finderImpl) FindAll() ([]ServiceOutput, error) {
	s, err := u.repo.FindAll()

	if err != nil {
		return []ServiceOutput{}, err
	}

	return xslices.Map(s, toServiceOutput), nil
}

func (u *finderImpl) FindByID(id string) (ServiceOutput, error) {
	s, err := u.repo.FindByID(nanoid.ID(id))

	if err != nil {
		return ServiceOutput{}, err
	}

	return toServiceOutput(s), err
}

func (u *finderImpl) FindByIDs(ids []string) ([]ServiceOutput, error) {
	_ids := xslices.Map(ids, shared.StringToNanoid)

	s, err := u.repo.FindByIDs(_ids)

	if err != nil {
		return []ServiceOutput{}, err
	}

	return xslices.Map(s, toServiceOutput), nil
}
