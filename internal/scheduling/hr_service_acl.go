package scheduling

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
)

var (
	ErrProfessionalNotFound = errors.New("professional not found")
)

type ProfessionalsACL interface {
	FindProfessionalByID(id nanoid.ID) (Professional, error)
}

type ProfessionalsACLFunc func(id nanoid.ID) (Professional, error)

func (f ProfessionalsACLFunc) FindProfessionalByID(id nanoid.ID) (Professional, error) {
	return f(id)
}
