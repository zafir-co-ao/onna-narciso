package inmem

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/sessions"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
)

type inmemSessionRepositoryImpl struct {
	shared.BaseRepository[sessions.Session]
}

func NewSessionRepository(s ...sessions.Session) sessions.Repository {
	return &inmemSessionRepositoryImpl{
		BaseRepository: shared.NewBaseRepository[sessions.Session](s...),
	}
}

func (s *inmemSessionRepositoryImpl) FindByID(id nanoid.ID) (sessions.Session, error) {
	for _, session := range s.Data {
		if session.ID.String() == id.String() {
			return session, nil
		}
	}
	return sessions.Session{}, sessions.ErrSessionNotFound
}

func (s *inmemSessionRepositoryImpl) FindByAppointmentsIDs(ids []nanoid.ID) ([]sessions.Session, error) {
	sessions := make([]sessions.Session, 0)

	for _, session := range s.Data {
		for _, id := range ids {
			if session.AppointmentID.String() == id.String() {
				sessions = append(sessions, session)
			}
		}
	}

	return sessions, nil
}

func (s *inmemSessionRepositoryImpl) Save(session sessions.Session) error {
	s.Data[session.ID] = session
	return nil
}
