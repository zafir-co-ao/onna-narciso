package scheduling

type WeeklyAppointmentsGetter interface {
	Get(week string, serviceID string, professionalsIDs []string) ([]Appointment, error)
}

type weeklyAppointmentsGetterImpl struct {
	repo AppointmentRepository
}

func NewWeeklyAppointmentsGetter(repo AppointmentRepository) WeeklyAppointmentsGetter {
	return &weeklyAppointmentsGetterImpl{repo: repo}
}

func (w *weeklyAppointmentsGetterImpl) Get(week string, serviceID string, professionalsIDs []string) ([]Appointment, error) {
	return w.repo.FindByWeekServiceAndProfessionals(week, serviceID, professionalsIDs)
}
