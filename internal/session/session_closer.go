package session

import "github.com/zafir-co-ao/onna-narciso/internal/shared/id"

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
	s, _ := u.repo.FindByID(id.NewID(i))

	s.Close()

	u.repo.Save(s)

	return nil
}
