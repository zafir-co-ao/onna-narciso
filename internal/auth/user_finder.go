package auth

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
)

type UserFinder interface {
	Find(id string) ([]UserOutput, error)
}

type finderImpl struct {
	repo Repository
}

func NewUserFinder(repo Repository) UserFinder {
	return &finderImpl{repo}
}

func (u *finderImpl) Find(id string) ([]UserOutput, error) {
	au, err := u.repo.FindByID(nanoid.ID(id))
	if err != nil {
		return []UserOutput{}, err
	}

	if !au.IsManager() {
		return []UserOutput{}, ErrUserNotAllowed
	}

	users, err := u.repo.FindAll()

	if err != nil {
		return []UserOutput{}, err
	}

	return xslices.Map(users, toUserOutput), nil
}
