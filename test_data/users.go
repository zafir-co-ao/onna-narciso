package testdata

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/auth"
)

var Users = []auth.User{
	{
		ID:          nanoid.ID("1"),
		Username:    "admin",
		Email:       "admin@gmail.com",
		PhoneNumber: "923452312",
		Password:    auth.MustNewPassword("admin@1234"),
		Role:        auth.RoleManager,
	},
	{
		ID:          nanoid.ID("2"),
		Username:    "john.doe",
		Email:       "john234@outlook.com",
		PhoneNumber: "934123456",
		Password:    auth.MustNewPassword("1234"),
		Role:        auth.RoleReceptionist,
	},
}
