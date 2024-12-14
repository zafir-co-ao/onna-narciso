package auth

import "errors"

var ErrAuthenticationFailed = errors.New("authentication failed")

type UserAuthenticator interface {
	Authenticate(i UserAuthenticatorInput) (UserOutput, error)
}

type authenticatorImpl struct {
	repo Repository
}

type UserAuthenticatorInput struct {
	Username string
	Password string
}

func NewUserAuthenticator(repo Repository) UserAuthenticator {
	return &authenticatorImpl{repo}
}

func (u *authenticatorImpl) Authenticate(i UserAuthenticatorInput) (UserOutput, error) {
	user, err := u.repo.FindByUsername(Username(i.Username))
	if err != nil {
		return UserOutput{}, ErrAuthenticationFailed
	}

	if !user.VerifyPassword(i.Password) {
		return UserOutput{}, ErrAuthenticationFailed
	}

	return toUserOutput(user), nil
}
