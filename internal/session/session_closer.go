package session

import (
	"github.com/zafir-co-ao/onna-narciso/internal/shared/event"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/id"
)

const EventSessionClosed = "EventSessionClosed"

type SessionCloserInput struct {
	SessionID   string
	ServicesIDs []string
}

type SessionCloser interface {
	Close(i SessionCloserInput) error
}

type sessionCloserImpl struct {
	repo SessionRepository
	sacl ServiceAcl
	bus  event.Bus
}

func NewSessionCloser(r SessionRepository, sacl ServiceAcl, b event.Bus) SessionCloser {
	return &sessionCloserImpl{repo: r, sacl: sacl, bus: b}
}

func (u *sessionCloserImpl) Close(i SessionCloserInput) error {
	s, err := u.repo.FindByID(id.NewID(i.SessionID))
	if err != nil {
		return ErrSessionNotFound
	}

	ids := id.ParseToIDs(i.ServicesIDs)

	services, err := u.sacl.FindByIDs(ids)
	if err != nil {
		return err
	}

	err = s.Close(services)
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
