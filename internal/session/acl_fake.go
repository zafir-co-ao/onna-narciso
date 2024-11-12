package session

import "github.com/kindalus/godx/pkg/nanoid"

type FakeServiceACL struct{}

func (f FakeServiceACL) FindByIDs(ids []nanoid.ID) ([]Service, error) {
	return EmptyServices, nil
}
