package scheduling

type AppointmentFinder interface {
	Execute(id string) (AppointmentOutput, error)
}

type appointmentFinderImpl struct {
	repo AppointmentRepository
}

func NewAppointmentFinder(r AppointmentRepository) AppointmentFinder {
	return &appointmentFinderImpl{repo: r}
}

func (f *appointmentFinderImpl) Execute(id string) (AppointmentOutput, error) {
	a, err := f.repo.FindByID(NewID(id))
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	return buildOutput(a), nil
}
