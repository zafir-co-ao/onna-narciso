package services

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/duration"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

type Service struct {
	ID          nanoid.ID
	Name        name.Name
	Duration    duration.Duration
	Description string
}

func NewService(name name.Name, duration duration.Duration, description string) Service {
	return Service{
		ID:          nanoid.New(),
		Name:        name,
		Duration:    duration,
		Description: description,
	}
}

func (s Service) GetID() nanoid.ID {
	return s.ID
}
