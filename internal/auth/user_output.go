package auth

type UserOutput struct {
	ID       string
	Username string
	Password string
}

func toUserOutput(u User) UserOutput {
	return UserOutput{
		Username: u.Username.String(),
		Password: u.Password.String(),
	}
}
