package sessions

import (
	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
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
	ids := xslices.Map(appointmentIDs, shared.StringToNanoid)

	s, err := u.repo.FindByAppointmentsIDs(ids)

	if err != nil {
		return []SessionOutput{}, err
	}

	return xslices.Map(s, toSessionOutput), nil
}
