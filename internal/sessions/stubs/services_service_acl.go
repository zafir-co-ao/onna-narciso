package stubs

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions"
)

var services = map[string]sessions.SessionService{
	"1": {ServiceID: nanoid.ID("1"), ProfessionalID: nanoid.ID("1")},
	"2": {ServiceID: nanoid.ID("2"), ProfessionalID: nanoid.ID("1")},
	"3": {ServiceID: nanoid.ID("3"), ProfessionalID: nanoid.ID("1")},
}

func NewServicesServiceACL() sessions.ServicesServiceACL {
	var f sessions.ServicesACLFunc = func(ids []nanoid.ID) ([]sessions.SessionService, error) {
		selected := make([]sessions.SessionService, 0)
		for _, id := range ids {
			if services[id.String()].ServiceID != id {
				return sessions.EmptyServices, sessions.ErrServiceNotFound
			}
			selected = append(selected, services[id.String()])
		}

		return selected, nil
	}

	return sessions.ServicesACLFunc(f)
}
