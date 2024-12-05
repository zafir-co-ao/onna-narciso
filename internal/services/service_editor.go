package services

import (
	"fmt"

	"github.com/kindalus/godx/pkg/nanoid"
)

type ServiceEditor interface {
	Edit(id string) error
}

type ServiceEditorImpl struct {
	repo Repository
}

func NewServiceEditor(repo Repository) ServiceEditor {
	return &ServiceEditorImpl{repo: repo}
}

func (u *ServiceEditorImpl) Edit(i string) error {
	fmt.Print("O ID: ", i)

	_, err := u.repo.FindByID(nanoid.ID(i))

	if err != nil {
		return err
	}

	return nil
}
