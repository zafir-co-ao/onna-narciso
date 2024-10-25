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

func (u *appointmentCancelerImpl) Execute(id string) error {
	a, err := u.repo.FindByID(NewID(id))
	if err != nil {
		return err
	}

	err = a.Cancel()
	if err != nil {
		return err
	}

	u.repo.Save(a)

	return nil
}
