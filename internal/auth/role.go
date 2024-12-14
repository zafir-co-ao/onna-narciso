package auth

import "errors"

const (
	RoleManager      Role = "Gestor"
	RoleCustomer     Role = "Cliente"
	RoleReceptionist Role = "Recepcionista"
)

var ErrRoleNotAllowed = errors.New("role of user not allowed")

var roles = map[string]Role{
	RoleManager.String():      RoleManager,
	RoleCustomer.String():     RoleCustomer,
	RoleReceptionist.String(): RoleReceptionist,
}

type Role string

func NewRole(v string) (Role, error) {
	if r, ok := roles[v]; ok {
		return r, nil
	}

	return Role(""), ErrRoleNotAllowed

}

func (r Role) String() string {
	return string(r)
}
