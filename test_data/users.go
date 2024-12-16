package testdata

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/auth"
)

var Users = []auth.User{
	{
		ID:       nanoid.ID("1"),
		Username: "admin",
		Password: auth.MustNewPassword("admin@1234"),
		Role:     auth.RoleManager,
	},
	{
		ID:       nanoid.ID("2"),
		Username: "john.doe",
		Password: auth.MustNewPassword("john.doe123"),
		Role:     auth.RoleReceptionist,
	},
}
