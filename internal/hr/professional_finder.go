package hr

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
)

type ProfessionalFinder interface {
	FindByID(id string) (ProfessionalOutput, error)
	FindAll() ([]ProfessionalOutput, error)
}

type finderImpl struct {
	repo Repository
}

func NewProfessionalFinder(repo Repository) ProfessionalFinder {
	return &finderImpl{repo}
}

func (u *finderImpl) FindByID(id string) (ProfessionalOutput, error) {
	p, err := u.repo.FindByID(nanoid.ID(id))
	if err != nil {
		return ProfessionalOutput{}, err
	}

	return toProfessionalOutput(p), nil
}

func (u *finderImpl) FindAll() ([]ProfessionalOutput, error) {
	professionals, err := u.repo.FindAll()

	if err != nil {
		return []ProfessionalOutput{}, err
	}

	return xslices.Map(professionals, toProfessionalOutput), nil
}
