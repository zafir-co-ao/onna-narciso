package sessions

import "github.com/kindalus/godx/pkg/nanoid"

type FakeServiceACL struct{}

func (f FakeServiceACL) FindByIDs(ids []nanoid.ID) ([]SessionService, error) {
	return EmptyServices, nil
}
