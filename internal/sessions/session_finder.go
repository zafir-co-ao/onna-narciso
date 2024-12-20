package sessions

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
)

type SessionFinder interface {
	Find(appointmentIDs []string) ([]SessionOutput, error)
	FindByID(id string) (SessionOutput, error)
}

type finderImpl struct {
	repo Repository
}

func NewSessionFinder(repo Repository) SessionFinder {
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

func (u *finderImpl) FindByID(id string) (SessionOutput, error) {
	s, err := u.repo.FindByID(nanoid.ID(id))
	if err != nil {
		return SessionOutput{}, err
	}

	return toSessionOutput(s), nil
}
