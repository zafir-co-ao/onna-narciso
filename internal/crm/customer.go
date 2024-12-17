package crm

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

const MinimumAgeAllowed = 12

var (
	ErrAgeNotAllowed          = errors.New("age not allowed")
	ErrNifAlreadyUsed         = errors.New("nif already used")
	ErrEmailAlreadyUsed       = errors.New("email already used")
	ErrPhoneNumberAlreadyUsed = errors.New("phone number already used")
)

type Customer struct {
	ID          nanoid.ID
	Name        name.Name
	Nif         Nif
	BirthDate   date.Date
	Email       Email
	PhoneNumber PhoneNumber
}

func (c Customer) GetID() nanoid.ID {
	return c.ID
}

func isAllowedAge(d date.Date) bool {
	return date.Today().Year()-d.Year() >= MinimumAgeAllowed
}

func getBirthDate(v string) (date.Date, error) {
	if len(v) == 0 {
		return date.Date(""), nil
	}

	return date.New(v)
}
