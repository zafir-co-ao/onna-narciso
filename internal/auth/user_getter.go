package auth

import "github.com/kindalus/godx/pkg/nanoid"

type UserGetter interface {
	Get(id string) (UserOutput, error)
}

type getterImpl struct {
	repo Repository
}

func NewUserGetter(repo Repository) UserGetter {
	return &getterImpl{repo}
}

func (u *getterImpl) Get(id string) (UserOutput, error) {
	user, err := u.repo.FindByID(nanoid.ID(id))

	if err != nil {
		return UserOutput{}, err
	}

	return toUserOutput(user), nil
}
