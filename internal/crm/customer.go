package crm

import (
	"errors"
	"time"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

const MinimumAgeAllowed = 12

var ErrAgeNotAllowed = errors.New("Age is less than 12 not allowed")

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

func isAllowedAge(d date.Date) bool {
	t, _ := time.Parse("2006-01-02", d.String())
	age := time.Now().Year() - t.Year()
	return age >= MinimumAgeAllowed
}
