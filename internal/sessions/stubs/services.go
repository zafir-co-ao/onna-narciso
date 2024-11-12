package stubs

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions"
)

var services = map[string]sessions.SessionService{
	"1": sessions.SessionService{ServiceID: nanoid.ID("1"), ProfessionalID: nanoid.ID("1")},
	"2": sessions.SessionService{ServiceID: nanoid.ID("2"), ProfessionalID: nanoid.ID("1")},
	"3": sessions.SessionService{ServiceID: nanoid.ID("3"), ProfessionalID: nanoid.ID("1")},
}

type serviceACLStub struct{}

func NewServiceACL() sessions.ServiceACL {
	return serviceACLStub{}
}

func (f serviceACLStub) FindByIDs(ids []nanoid.ID) ([]sessions.SessionService, error) {
	selected := make([]sessions.SessionService, 0)
	for _, id := range ids {
		if services[id.String()].ServiceID.String() != id.String() {
			return sessions.EmptyServices, sessions.ErrServiceNotFound
		}
		selected = append(selected, services[id.String()])
	}
	return selected, nil
}
