package services

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/services/price"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/duration"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

type Description string

type Service struct {
	ID          nanoid.ID
	Name        name.Name
	Duration    duration.Duration
	Price       price.Price
	Description Description
}

func NewService(
	name name.Name,
	duration duration.Duration,
	price price.Price,
	description Description,
) Service {
	return Service{
		ID:          nanoid.New(),
		Name:        name,
		Duration:    duration,
		Description: description,
		Price:       price,
	}
}

func (s Service) GetID() nanoid.ID {
	return s.ID
}
