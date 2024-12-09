package testdata

import "github.com/zafir-co-ao/onna-narciso/internal/scheduling"

var Services []scheduling.Service = []scheduling.Service{
	{ID: "1", Name: "Manicure"},
	{ID: "2", Name: "Pedicure"},
	{ID: "3", Name: "Depilação"},
	{ID: "4", Name: "Massagem", Duration: 120},
}
