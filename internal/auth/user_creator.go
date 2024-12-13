package auth

import "github.com/kindalus/godx/pkg/nanoid"

type UserCreatorInput struct {
	Username string
	Password string
	Role     string
}

type UserCreator interface {
	Create(i UserCreatorInput) (UserOutput, error)
}

type creatorImpl struct {
	repo Repository
}

func NewUserCreator(repo Repository) UserCreator {
	return &creatorImpl{repo}
}

func (u *creatorImpl) Create(i UserCreatorInput) (UserOutput, error) {
	user := User{
		ID:       nanoid.New(),
		Username: Username(i.Username),
		Password: Password(i.Password),
		Role:     Role(i.Role),
	}

	err := u.repo.Save(user)
	if err != nil {
		return UserOutput{}, err
	}

	return UserOutput{ID: user.ID.String()}, nil
}
