package scheduling

import "errors"

var ErrProfessionalNotFound = errors.New("Professional not found")

type ProfessionalRepository interface {
	Get(id string) (Professional, error)
	Save(p Professional) error
}
