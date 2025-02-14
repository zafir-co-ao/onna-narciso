package auth

type UserOutput struct {
	ID          string
	Username    string
	Email       string
	PhoneNumber string
	Role        string
}

func toUserOutput(u User) UserOutput {
	return UserOutput{
		ID:          u.ID.String(),
		Username:    u.Username.String(),
		Email:       u.Email.String(),
		PhoneNumber: u.PhoneNumber.String(),
		Role:        u.Role.String(),
	}
}
