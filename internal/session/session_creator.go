package session

import "github.com/zafir-co-ao/onna-narciso/internal/shared/id"

type SessionOutput struct {
	ID            string
	AppointmentID string
}

type SessionCreator interface {
	Create(appointmentID string) (SessionOutput, error)
}

type sessionCreatorImpl struct {
	repo SessionRepository
}

func (c *sessionCreatorImpl) Create(appointmentID string) (SessionOutput, error) {

	_id, err := id.Random()
	if err != nil {
		return SessionOutput{}, err
	}

	s := Session{
		ID:            _id,
		AppointmentID: id.NewID(appointmentID),
	}

	c.repo.Save(s)

	return SessionOutput{
		ID:            s.ID.String(),
		AppointmentID: s.AppointmentID.String(),
	}, nil
}

func NewSessionCreator(r SessionRepository) SessionCreator {
	return &sessionCreatorImpl{repo: r}
}
