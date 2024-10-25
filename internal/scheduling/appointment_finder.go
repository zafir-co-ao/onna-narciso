package scheduling

type AppointmentFinder interface {
	Execute(id string) (Appointment, error)
}

type appointmentFinderImpl struct {
	repo AppointmentRepository
}

func NewAppointmentFinder(r AppointmentRepository) AppointmentFinder {
	return &appointmentFinderImpl{repo: r}
}

func (f *appointmentFinderImpl) Execute(id string) (Appointment, error) {
	a, _ := f.repo.FindByID(NewID(id))
	return a, nil
}
