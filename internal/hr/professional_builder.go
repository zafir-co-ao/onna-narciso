package hr

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

type professionalBuilder Professional

func NewProfessionalBuilder() *professionalBuilder {
	return &professionalBuilder{ID: nanoid.New()}
}

func (p *professionalBuilder) WithID(id nanoid.ID) *professionalBuilder {
	p.ID = id
	return p
}

func (p *professionalBuilder) WithName(n name.Name) *professionalBuilder {
	p.Name = n
	return p
}

func (p *professionalBuilder) WithServices(s []Service) *professionalBuilder {
	p.Services = s
	return p
}

func (p professionalBuilder) Build() Professional {
	return Professional(p)
}
