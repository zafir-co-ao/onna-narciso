package sessions

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
)

type inmemSessionRepositoryImpl struct {
	shared.BaseRepository[Session]
}

func NewInmemRepository(s ...Session) Repository {
	return &inmemSessionRepositoryImpl{BaseRepository: shared.NewBaseRepository[Session](s...)}
}

func (r *inmemSessionRepositoryImpl) FindByID(id nanoid.ID) (Session, error) {
	for _, s := range r.Data {
		if s.ID.String() == id.String() {
			return s, nil
		}
	}
	return Session{}, ErrSessionNotFound
}

func (r *inmemSessionRepositoryImpl) FindByAppointmentsIDs(ids []nanoid.ID) ([]Session, error) {
	sessions := make([]Session, 0)

	for _, s := range r.Data {
		for _, id := range ids {
			if s.AppointmentID.String() == id.String() {
				sessions = append(sessions, s)
			}
		}
	}

	return sessions, nil
}

func (r *inmemSessionRepositoryImpl) Save(session Session) error {
	r.Data[session.ID] = session
	return nil
}
