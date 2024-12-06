package crm

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

type CustomerEditor interface {
	Edit(id, newName string) error
}

type CustomerEditorImpl struct {
	repo Repository
}

func NewCustomerEditor(repo Repository) CustomerEditor {
	return &CustomerEditorImpl{repo: repo}
}

func (u *CustomerEditorImpl) Edit(id, newName string) error {
	_, err := u.repo.FindByID(nanoid.ID(id))
	if err != nil {
		return err
	}

	n, err := name.New(newName)
	if err != nil {
		return err
	}

	c := NewCustomerBuilder().
		WithName(n).Build()

	err = u.repo.Save(c)

	return nil
}
