package services

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/duration"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

type serviceBuilder Service

func NewServiceBuilder() *serviceBuilder {
	return &serviceBuilder{}
}

func (s *serviceBuilder) WithID(id nanoid.ID) *serviceBuilder {
	s.ID = id
	return s
}

func (s *serviceBuilder) WithName(n name.Name) *serviceBuilder {
	s.Name = n
	return s
}

func (s *serviceBuilder) WithDescription(d Description) *serviceBuilder {
	s.Description = d
	return s
}

func (s *serviceBuilder) WithPrice(p Price) *serviceBuilder {
	s.Price = p
	return s
}

func (s *serviceBuilder) WithDiscount(d Discount) *serviceBuilder {
	s.Discount = d
	return s
}

func (s *serviceBuilder) WithDuration(d duration.Duration) *serviceBuilder {
	s.Duration = d
	return s
}

func (s serviceBuilder) Build() Service {
	return Service(s)
}
