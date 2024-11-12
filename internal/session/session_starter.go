package session

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
)

const EventSessionStarted = "EventSessionStarted"

type Starter interface {
	Start(id string) error
}

type starterImpl struct {
	repo Repository
	bus  event.Bus
}

func NewSessionStarter(repo Repository, bus event.Bus) Starter {
	return &starterImpl{repo, bus}
}

func (u *starterImpl) Start(id string) error {
	s, err := u.repo.FindByID(nanoid.ID(id))
	if err != nil {
		return ErrSessionNotFound
	}

	err = s.Start()
	if err != nil {
		return err
	}

	err = u.repo.Save(s)
	if err != nil {
		return err
	}

	e := event.New(
		EventSessionStarted,
		event.WithHeader(event.HeaderAggregateID, id),
		event.WithPayload(s),
	)

	u.bus.Publish(e)

	return nil
}
