package components

type link struct {
	title string
	url   string
}

var menu = []link{
	{
		title: "Agenda",
		url:   "/daily-appointments",
	},
	{
		title: "Serviços",
		url:   "/services",
	},
	{
		title: "Clientes",
		url:   "/customers",
	},
	{
		title: "Profissionais",
		url:   "/professionals",
	},
	{
		title: "Utilizadores",
		url:   "/auth/users",
	},
}
