package inmem

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/session"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
)

type inmemeSessionRepositoryImpl struct {
	shared.BaseRepository[session.Session]
}

func (s *inmemeSessionRepositoryImpl) FindByID(id nanoid.ID) (session.Session, error) {
	for _, session := range s.Data {
		if session.ID.String() == id.String() {
			return session, nil
		}
	}
	return session.Session{}, session.ErrSessionNotFound
}

func NewSessionRepository(s ...session.Session) session.Repository {
	return &inmemeSessionRepositoryImpl{
		BaseRepository: shared.NewBaseRepository[session.Session](s...),
	}
}

func (s *inmemeSessionRepositoryImpl) Save(session session.Session) error {
	s.Data[session.ID] = session
	return nil
}
