package services

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
)

var ErrServiceNotFound = errors.New("service not found")

type Repository interface{
	FindByID(id nanoid.ID) (Service, error)
	Save(s Service) error
}
