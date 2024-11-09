package session

import (
	"github.com/zafir-co-ao/onna-narciso/internal/shared/event"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/id"
)

const EventSessionClosed = "EventSessionClosed"

type SessionCloser interface {
	Close(i string) error
}

type sessionCloserImpl struct {
	repo SessionRepository
	bus  event.Bus
}

func NewSessionCloser(r SessionRepository, b event.Bus) SessionCloser {
	return &sessionCloserImpl{repo: r, bus: b}
}

func (u *sessionCloserImpl) Close(i string) error {
	s, err := u.repo.FindByID(id.NewID(i))
	if err != nil {
		return ErrSessionNotFound
	}

	err = s.Close()
	if err != nil {
		return err
	}

	err = u.repo.Save(s)
	if err != nil {
		return err
	}

	u.bus.Publish(event.New(EventSessionClosed))

	return nil
}
