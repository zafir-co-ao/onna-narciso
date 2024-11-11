package inmem

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/session"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
)

type inmemSessionRepositoryImpl struct {
	shared.BaseRepository[session.Session]
}

func NewSessionRepository(s ...session.Session) session.Repository {
	return &inmemSessionRepositoryImpl{
		BaseRepository: shared.NewBaseRepository[session.Session](s...),
	}
}

func (s *inmemSessionRepositoryImpl) FindByID(id nanoid.ID) (session.Session, error) {
	for _, session := range s.Data {
		if session.ID.String() == id.String() {
			return session, nil
		}
	}
	return session.Session{}, session.ErrSessionNotFound
}

func (s *inmemSessionRepositoryImpl) FindByAppointmentsIDs(ids []nanoid.ID) ([]session.Session, error) {
	sessions := make([]session.Session, 0)

	for _, session := range s.Data {
		for _, id := range ids {
			if session.AppointmentID.String() == id.String() {
				sessions = append(sessions, session)
			}
		}
	}

	return sessions, nil
}

func (s *inmemSessionRepositoryImpl) Save(session session.Session) error {
	s.Data[session.ID] = session
	return nil
}
