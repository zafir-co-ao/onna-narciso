package hr

import (
	"slices"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

var EventProfessionalCreated = "EventProfessionalCreated"

type Service struct {
	ID   nanoid.ID
	Name name.Name
}

type Professional struct {
	ID       nanoid.ID
	Name     name.Name
	Services []Service
}

func (p *Professional) HasService(id nanoid.ID) bool {
	return slices.ContainsFunc(p.Services, func(s Service) bool {
		return s.ID == id
	})
}

func (p Professional) GetID() nanoid.ID {
	return p.ID
}
