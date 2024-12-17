package sessions

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
)

type Finder interface {
	Find(appointmentIDs []string) ([]SessionOutput, error)
}

type finderImpl struct {
	repo Repository
}

func NewSessionFinder(repo Repository) Finder {
	return &finderImpl{repo}
}

func (u *finderImpl) Find(appointmentIDs []string) ([]SessionOutput, error) {
	ids := xslices.Map(appointmentIDs, func(id string) nanoid.ID { return nanoid.ID(id) })

	s, err := u.repo.FindByAppointmentsIDs(ids)

	if err != nil {
		return []SessionOutput{}, err
	}

	return xslices.Map(s, toSessionOutput), nil
}
