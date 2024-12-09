package crm

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

type Customer struct {
	ID          nanoid.ID
	Name        name.Name
	Nif         Nif
	BirthDate   date.Date
	Email       Email
	PhoneNumber PhoneNumber
}

func (c *Customer) IsSameNif(n Nif) bool {
	return c.Nif.String() == n.String()
}

func (c Customer) GetID() nanoid.ID {
	return c.ID
}
