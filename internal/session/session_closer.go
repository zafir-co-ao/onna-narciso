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

func NewSessionCloser(repo SessionRepository, sacl ServiceAcl, bus event.Bus) SessionCloser {
	return &sessionCloserImpl{repo, sacl, bus}
}

func (u *sessionCloserImpl) Close(i SessionCloserInput) error {
	s, err := u.repo.FindByID(id.NewID(i.SessionID))
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

func (u *sessionCloserImpl) findServices(ids []string) ([]Service, error) {
	if len(ids) == 0 {
		return EmptyServices, nil
	}

	s, err := u.sacl.FindByIDs(id.ParseToIDs(ids))

	if err != nil {
		return EmptyServices, err
	}

	return s, nil
}
