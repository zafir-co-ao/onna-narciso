package crm

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/email"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/nif"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/phone"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

type Customer struct {
	ID          nanoid.ID
	Name        name.Name
	Nif         nif.Nif
	BirthDate   date.Date
	Email       email.Email
	PhoneNumber phone.PhoneNumber
}

func (c *Customer) IsSameNif(n nif.Nif) bool {
	return c.Nif.String() == n.String()
}

func (c Customer) GetID() nanoid.ID {
	return c.ID
}
