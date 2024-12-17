package hr

import "github.com/kindalus/godx/pkg/nanoid"

type ProfessionalOutput struct {
	ID   string
	Name string
}

type ProfessionalFinder interface {
	FindByID(id string) (ProfessionalOutput, error)
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

	return ProfessionalOutput{
		ID:   p.ID.String(),
		Name: p.Name.String(),
	}, nil
}
