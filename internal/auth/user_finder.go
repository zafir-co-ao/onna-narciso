package auth

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
)

type UserFinder interface {
	FindAll() ([]UserOutput, error)
	FindByID(id string) (UserOutput, error)
}

type finderImpl struct {
	repo Repository
}

func NewUserFinder(repo Repository) UserFinder {
	return &finderImpl{repo}
}

func (u *finderImpl) FindAll() ([]UserOutput, error) {
	users, err := u.repo.FindAll()

	if err != nil {
		return []UserOutput{}, err
	}

	return xslices.Map(users, toUserOutput), nil
}

func (u *finderImpl) FindByID(id string) (UserOutput, error) {
	user, err := u.repo.FindByID(nanoid.ID(id))

	if err != nil {
		return UserOutput{}, err
	}

	return toUserOutput(user), nil
}
