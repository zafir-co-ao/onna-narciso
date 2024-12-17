package sessions

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
)

const EventSessionClosed = "EventSessionClosed"

type CloserInput struct {
	SessionID   string
	ServicesIDs []string
}

type Closer interface {
	Close(i CloserInput) error
}

type closerImpl struct {
	repo Repository
	sacl ServicesServiceACL
	bus  event.Bus
}

func NewSessionCloser(repo Repository, sacl ServicesServiceACL, bus event.Bus) Closer {
	return &closerImpl{repo, sacl, bus}
}

func (u *closerImpl) Close(i CloserInput) error {
	s, err := u.repo.FindByID(nanoid.ID(i.SessionID))
	if err != nil {
		return ErrSessionNotFound
	}

	svc, err := u.findServices(i.ServicesIDs)
	if err != nil {
		return err
	}

	err = s.Close(svc)
	if err != nil {
		return err
	}

	err = u.repo.Save(s)
	if err != nil {
		return err
	}

	e := event.New(
		EventSessionClosed,
		event.WithHeader(event.HeaderAggregateID, s.ID.String()),
		event.WithPayload(i),
	)

	u.bus.Publish(e)

	return nil
}

func (u *closerImpl) findServices(ids []string) ([]SessionService, error) {
	if len(ids) == 0 {
		return EmptyServices, nil
	}

	_ids := xslices.Map(ids, func(id string) nanoid.ID { return nanoid.ID(id) })

	return u.sacl.FindByIDs(_ids)
}
