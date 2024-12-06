package crm

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/nif"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

type CustomerEditorInput struct {
	ID        nanoid.ID
	Name      string
	NIF       string
	BirthDate string
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

	c := NewCustomerBuilder().
		WithID(i.ID).
		WithName(n).
		WithNif(nif).
		WithBirthDate(date.Date(i.BirthDate)).
		Build()

	err = u.repo.Save(c)

	return nil
}
