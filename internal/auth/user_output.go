package auth

type UserOutput struct {
	ID       string
	Username string
	Role     string
}

func toUserOutput(u User) UserOutput {
	return UserOutput{
		ID:       u.ID.String(),
		Username: u.Username.String(),
		Role:     u.Role.String(),
	}
}
