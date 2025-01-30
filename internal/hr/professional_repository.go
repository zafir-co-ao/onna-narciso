package hr

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
)

var ErrProfessionalNotFound = errors.New("professional not found")

type Repository interface {
	FindByID(id nanoid.ID) (Professional, error)
	FindAll() ([]Professional, error)
	Save(p Professional) error
}
