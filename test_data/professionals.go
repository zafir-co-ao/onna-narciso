package testdata

import (
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/id"
)

var Professionals []scheduling.Professional = []scheduling.Professional{
	{ID: "1", Name: "Joana DArc", ServicesIDS: []id.ID{"1", "2"}},
	{ID: "2", Name: "Napoléon Bonaparte", ServicesIDS: []id.ID{"3"}},
	{ID: "3", Name: "Alexandre III", ServicesIDS: []id.ID{"4"}},
	{ID: "4", Name: "Cleopatra", ServicesIDS: []id.ID{"1", "2", "3", "4"}},
}
