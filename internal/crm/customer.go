package crm

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

type Nif string

func (n Nif) String() string {
	return string(n)
}

type Customer struct {
	ID          nanoid.ID
	Name        name.Name
	Nif         Nif
	BirthDate   date.Date
	Email       string
	PhoneNumber string
}

func (c Customer) GetID() nanoid.ID {
	return c.ID
}
