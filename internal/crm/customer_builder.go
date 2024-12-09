package crm

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

type customerBuilder Customer

func NewCustomerBuilder() *customerBuilder {
	return &customerBuilder{ID: nanoid.New()}
}

func (c *customerBuilder) WithName(n name.Name) *customerBuilder {
	c.Name = n
	return c
}

func (c *customerBuilder) WithID(id nanoid.ID) *customerBuilder {
	c.ID = id
	return c
}

func (c *customerBuilder) WithNif(nif Nif) *customerBuilder {
	c.Nif = nif
	return c
}

func (c *customerBuilder) WithEmail(e Email) *customerBuilder {
	c.Email = e
	return c
}

func (c *customerBuilder) WithPhoneNumber(p PhoneNumber) *customerBuilder {
	c.PhoneNumber = p
	return c
}

func (c *customerBuilder) WithBirthDate(d date.Date) *customerBuilder {
	c.BirthDate = d
	return c
}

func (c customerBuilder) Build() Customer {
	return Customer(c)
}
