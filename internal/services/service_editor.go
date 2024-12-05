package services

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/services/price"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

type ServiceEditor interface {
	Edit(id nanoid.ID, name, newPrice string) error
}

type ServiceEditorImpl struct {
	repo Repository
}

func NewServiceEditor(repo Repository) ServiceEditor {
	return &ServiceEditorImpl{repo: repo}
}

func (u *ServiceEditorImpl) Edit(i nanoid.ID, newName string, newPrice string) error {

	_, err := u.repo.FindByID(i)
	if err != nil {
		return err
	}

	_name, err := name.New(newName)
	if err != nil {
		return err
	}

	_price, err := price.New(newPrice)

	s := NewServiceBuilder().WithID(i).
		WithName(_name).
		WithPrice(_price).
		Build()

	err = u.repo.Save(s)
	if err != nil {
		return err
	}

	return nil
}
