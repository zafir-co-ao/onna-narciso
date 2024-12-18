package stubs

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/hr"
)

var Services = map[string]hr.Service{
	"1": {ID: nanoid.ID("1")},
	"2": {ID: nanoid.ID("2")},
}

func NewServicesServiceACL() hr.ServicesServiceACL {
	var f hr.ServicesServiceACLFunc = func(ids []nanoid.ID) ([]hr.Service, error) {
		var services []hr.Service

		for _, id := range ids {
			if Services[id.String()].ID != id {
				return []hr.Service{}, hr.ErrServiceNotFound
			}

			services = append(services, Services[id.String()])
		}

		return services, nil
	}
	return f
}
