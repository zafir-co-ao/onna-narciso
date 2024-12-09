package sessions

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
)

type inmemSessionRepositoryImpl struct {
	shared.BaseRepository[Session]
}

func NewInmemRepository(s ...Session) Repository {
	return &inmemSessionRepositoryImpl{
		BaseRepository: shared.NewBaseRepository[Session](s...),
	}
}

func (s *inmemSessionRepositoryImpl) FindByID(id nanoid.ID) (Session, error) {
	for _, session := range s.Data {
		if session.ID.String() == id.String() {
			return session, nil
		}
	}
	return Session{}, ErrSessionNotFound
}

func (s *inmemSessionRepositoryImpl) FindByAppointmentsIDs(ids []nanoid.ID) ([]Session, error) {
	sessions := make([]Session, 0)

	for _, session := range s.Data {
		for _, id := range ids {
			if session.AppointmentID.String() == id.String() {
				sessions = append(sessions, session)
			}
		}
	}

	return sessions, nil
}

func (s *inmemSessionRepositoryImpl) Save(session Session) error {
	s.Data[session.ID] = session
	return nil
}
