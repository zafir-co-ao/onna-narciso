package components

import "github.com/zafir-co-ao/onna-narciso/internal/auth"

type link struct {
	title string
	url   string
}

var menus = map[string][]link{
	auth.RoleManager.String(): {
		{"Agenda", "/daily-appointments"},
		{"Serviços", "/services"},
		{"Clientes", "/customers"},
		{"Profissionais", "/professionals"},
		{"Utilizadores", "/auth/users"},
	},
	auth.RoleReceptionist.String(): {
		{"Agenda", "/daily-appointments"},
		{"Clientes", "/customers"},
	},
}

func GetMenu(role string) []link {
	return menus[role]
}
