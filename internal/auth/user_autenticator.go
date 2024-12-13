package auth

import "errors"

type UserAutenticator interface {
	Autenticate(i UserAutenticatorInput) (UserOutput, error)
}

type userAutenticatorImpl struct {
	repo Repository
}

type UserAutenticatorInput struct {
	Username string
	Password string
}

var ErrPasswordInvalid = errors.New("password invalid")

func NewUserAutenticator(repo Repository) UserAutenticator {
	return &userAutenticatorImpl{repo}
}

func (u *userAutenticatorImpl) Autenticate(i UserAutenticatorInput) (UserOutput, error) {

	user, err := u.repo.FindByUserName(Username(i.Username))
	if err != nil {
		return UserOutput{}, err
	}

	if !user.IsSamePassword(Password(i.Password)) {
		return UserOutput{}, ErrPasswordInvalid
	}

	return toUserOutput(user), nil
}
