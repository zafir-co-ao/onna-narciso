package components

import "github.com/zafir-co-ao/onna-narciso/internal/auth"

type UserUpdateParams struct {
	Url            string
	User           auth.UserOutput
	AuthUser       auth.UserOutput
	HxTarget       string
	HxTriggerEvent string
}
