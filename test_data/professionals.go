package testdata

import (
	"slices"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
)

var Professionals []scheduling.Professional = []scheduling.Professional{
	{ID: "1", Name: "Joana DArc", ServicesIDS: []nanoid.ID{"1", "2"}},
	{ID: "2", Name: "Napoléon Bonaparte", ServicesIDS: []nanoid.ID{"3"}},
	{ID: "3", Name: "Alexandre III", ServicesIDS: []nanoid.ID{"4"}},
	{ID: "4", Name: "Cleopatra", ServicesIDS: []nanoid.ID{"1", "2", "3", "4"}},
}

func FindProfessionalsByServiceID(serviceId string) []scheduling.Professional {
	professionals := make([]scheduling.Professional, 0)
	for _, p := range Professionals {
		if slices.Contains(p.ServicesIDS, nanoid.ID(serviceId)) {
			professionals = append(professionals, p)
		}
	}

	return professionals
}
