package sessions

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
)

type SessionOutput struct {
	ID            string
	AppointmentID string
	Status        string
}

type Finder interface {
	Find(apointmentIDs []string) ([]SessionOutput, error)
}

type finderImpl struct {
	repo Repository
}

func NewSessionFinder(r Repository) Finder {
	return &finderImpl{repo: r}
}

func (u *finderImpl) Find(appointmentIDs []string) ([]SessionOutput, error) {
	ids := xslices.Map(appointmentIDs, func(i string) nanoid.ID { return nanoid.ID(i) })

	s, err := u.repo.FindByAppointmentsIDs(ids)

	if err != nil {
		return []SessionOutput{}, err
	}

	return xslices.Map(s, toSessionOutput), nil
}

func toSessionOutput(s Session) SessionOutput {
	return SessionOutput{
		ID:            s.ID.String(),
		AppointmentID: s.AppointmentID.String(),
		Status:        string(s.Status),
	}
}
