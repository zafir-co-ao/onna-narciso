package scheduling

type DailyAppointmentsGetter interface {
	Get(day string) ([]Appointment, error)
}

type dailyAppointmentsGetterImpl struct {
	repo AppointmentRepository
}

func NewDailyAppointmentsGetter(repo AppointmentRepository) DailyAppointmentsGetter {
	return &dailyAppointmentsGetterImpl{repo: repo}
}

func (d *dailyAppointmentsGetterImpl) Get(date string) ([]Appointment, error) {
	return d.repo.FindByDate(Date(date))
}
