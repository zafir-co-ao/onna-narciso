package components

import "github.com/zafir-co-ao/onna-narciso/internal/auth"

type UserPasswordUpdateParams struct {
	Url            string
	User           auth.UserOutput
	HxTarget       string
	HxTriggerEvent string
}
