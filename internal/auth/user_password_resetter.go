package auth

type UserPasswordResetter interface {
	Reset(i UserPasswordResetterInput) error
}

type UserPasswordResetterInput struct {
	Email string
}

func NewUserPasswordResetter(r Repository) UserPasswordResetter {
	return &userPasswordResetterImpl{r}
}

type userPasswordResetterImpl struct {
	repo Repository
}

func (u *userPasswordResetterImpl) Reset(i UserPasswordResetterInput) error {
	user, err := u.repo.FindByEmail(Email(i.Email))
	if err != nil {
		return err
	}

	p, err := GeneratePassword(PasswordLength)
	if err != nil {
		return err
	}

	user.ResetPassword(p)

	err = u.repo.Save(user)
	if err != nil {
		return err
	}

	return nil
}
