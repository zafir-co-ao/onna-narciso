package auth

import "github.com/kindalus/godx/pkg/xslices"

func IsAvailableUsername(users []User, username Username) bool {
	return xslices.All(users, func(u User) bool {
		return u.Username != username
	})
}

func IsAvailableEmail(users []User, email Email) bool {
	return xslices.All(users, func(u User) bool {
		return u.Email != email
	})
}

func IsAvailablePhoneNumber(users []User, phoneNumber PhoneNumber) bool {
	return xslices.All(users, func(u User) bool {
		return u.PhoneNumber != phoneNumber
	})
}
