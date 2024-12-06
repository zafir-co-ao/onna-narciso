package crm

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/nif"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

type CustomerEditor interface {
	Edit(id, newName, newNif string) error
}

type CustomerEditorImpl struct {
	repo Repository
}

func NewCustomerEditor(repo Repository) CustomerEditor {
	return &CustomerEditorImpl{repo: repo}
}

func (u *CustomerEditorImpl) Edit(id, newName, newNif string) error {
	_, err := u.repo.FindByID(nanoid.ID(id))
	if err != nil {
		return err
	}

	n, err := name.New(newName)
	if err != nil {
		return err
	}

	nif, err := nif.New(newNif)

	c := NewCustomerBuilder().
		WithName(n).WithNif(nif).
		Build()

	err = u.repo.Save(c)

	return nil
}
