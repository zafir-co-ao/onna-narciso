package session

import (
	"github.com/zafir-co-ao/onna-narciso/internal/shared/id"
)

type SessionCloser interface {
	Close(i string) error
}

type sessionCloserImpl struct {
	repo SessionRepository
}

func NewSessionCloser(r SessionRepository) SessionCloser {
	return &sessionCloserImpl{repo: r}
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

	return nil
}
