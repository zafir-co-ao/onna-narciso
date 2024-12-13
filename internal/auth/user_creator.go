package auth

import "github.com/kindalus/godx/pkg/nanoid"

type UserOutput struct {
	ID string
}

type UserCreator interface {
	Create() (UserOutput, error)
}

type creatorImpl struct {
	repo Repository
}

func NewUserCreator(repo Repository) UserCreator {
	return &creatorImpl{repo}
}

func (u *creatorImpl) Create() (UserOutput, error) {
	user := User{ID: nanoid.New()}

	err := u.repo.Save(user)
	if err != nil {
		return UserOutput{}, err
	}

	return UserOutput{ID: user.ID.String()}, nil
}
