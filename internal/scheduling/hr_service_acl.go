package scheduling

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/hr"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

var ErrProfessionalNotFound = errors.New("professional not found")

type HRServiceACL interface {
	FindProfessionalByID(id nanoid.ID) (Professional, error)
}

type HRServiceACLFunc func(id nanoid.ID) (Professional, error)

func (f HRServiceACLFunc) FindProfessionalByID(id nanoid.ID) (Professional, error) {
	return f(id)
}

type internalHRServiceACL struct {
	finder hr.ProfessionalFinder
}

func NewInternalHRServiceACL(finder hr.ProfessionalFinder) HRServiceACL {
	return &internalHRServiceACL{finder}
}

func (i *internalHRServiceACL) FindProfessionalByID(id nanoid.ID) (Professional, error) {
	p, err := i.finder.FindByID(id.String())

	if err != nil {
		return Professional{}, ErrProfessionalNotFound
	}

	return Professional{
		ID:   nanoid.ID(p.ID),
		Name: name.Name(p.Name),
	}, nil
}
