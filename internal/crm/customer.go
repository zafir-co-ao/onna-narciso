package crm

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

const MinimumAgeAllowed = 12

var ErrAgeNotAllowed = errors.New("age not allowed")

type Customer struct {
	ID          nanoid.ID
	Name        name.Name
	Nif         Nif
	BirthDate   date.Date
	Email       Email
	PhoneNumber PhoneNumber
}

func (c *Customer) IsSameNif(n Nif) bool {
	return c.Nif == n
}

func (c *Customer) IsSameEmail(e Email) bool {
	return c.Email == e
}

func (c *Customer) IsSamePhoneNumber(p PhoneNumber) bool {
	return c.PhoneNumber == p
}

func (c Customer) GetID() nanoid.ID {
	return c.ID
}

func isAllowedAge(d date.Date) bool {
	return date.Today().Year()-d.Year() >= MinimumAgeAllowed
}
