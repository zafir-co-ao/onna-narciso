package scheduling

type AppointmentCanceler interface {
	Execute(id string) error
}

type appointmentCancelerImpl struct {
	repo AppointmentRepository
}

func NewAppointmentCanceler(repo AppointmentRepository) AppointmentCanceler {
	return &appointmentCancelerImpl{repo}
}

func (c *appointmentCancelerImpl) Execute(id string) error {
	a, err := c.repo.Get(id)
	if err != nil {
		return err
	}

	if a.IsCancelled() {
		return ErrInvalidStatusToCancel
	}

	a.Cancel()

	c.repo.Save(a)

	return nil
}
