package testdata

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
)

var Customers1 = []crm.Customer{
	{
		ID:        nanoid.ID("1"),
		Name:      "Paola Oliveira",
		Nif:       "004843261LA001",
		BirthDate: date.Today().AddDate(-12, 0, 0),
		Email:     "paola.oliveira@gmail.com",
	},
	{
		ID:        nanoid.ID("2"),
		Name:      "Juliana Paes",
		Nif:       "004843261LA002",
		BirthDate: date.Today().AddDate(-12, 0, 0),
	},
	{
		ID:        nanoid.ID("3"),
		Name:      "Gisele Bündchen",
		Nif:       "004843261LA003",
		BirthDate: date.Today().AddDate(-12, 0, 0),
	},
	{
		ID:        nanoid.ID("4"),
		Name:      "James Blunt",
		Nif:       "004843261LA004",
		BirthDate: date.Today().AddDate(-12, 0, 0),
		Email:     "james.blunt@gmail.com",
	},
}

var Services1 = []services.Service{
	{
		ID:       nanoid.New(),
		Name:     "Manicure",
		Price:    "1000",
		Duration: 45,
	},
	{
		ID:       nanoid.New(),
		Name:     "Pedicure",
		Price:    "4000",
		Duration: 60,
	},
	{
		ID:       nanoid.New(),
		Name:     "Depilação",
		Price:    "8000",
		Duration: 45,
	},
	{
		ID:       nanoid.New(),
		Name:     "Massagem",
		Price:    "11000",
		Duration: 90,
	},
}