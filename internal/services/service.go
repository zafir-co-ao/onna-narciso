package services

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/duration"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

type Description string

type Service struct {
	ID          nanoid.ID
	Name        name.Name
	Duration    duration.Duration
	Price       Price
	Discount    Discount
	Description Description
}

func (s Service) GetID() nanoid.ID {
	return s.ID
}

func (s Service) IsSamePrice(p Price) bool {
	return s.Price == p
}
