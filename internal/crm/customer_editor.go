package crm

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/email"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/nif"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

type CustomerEditorInput struct {
	ID        nanoid.ID
	Name      string
	NIF       string
	BirthDate string
	Email     string
}

type CustomerEditor interface {
	Edit(i CustomerEditorInput) error
}

type CustomerEditorImpl struct {
	repo Repository
}

func NewCustomerEditor(repo Repository) CustomerEditor {
	return &CustomerEditorImpl{repo: repo}
}

func (u *CustomerEditorImpl) Edit(i CustomerEditorInput) error {
	_, err := u.repo.FindByID(i.ID)
	if err != nil {
		return err
	}

	n, err := name.New(i.Name)
	if err != nil {
		return err
	}

	nif, err := nif.New(i.NIF)
	if err != nil {
		return err
	}

	if !date.IsValidFormat(i.BirthDate) {
		return date.ErrInvalidFormat
	}

	email, err := email.New(i.Email)
	if err != nil {
		return err
	}

	c := NewCustomerBuilder().
		WithID(i.ID).
		WithName(n).
		WithNif(nif).
		WithBirthDate(date.Date(i.BirthDate)).
		WithEmail(email).
		Build()

	err = u.repo.Save(c)

	return nil
}
