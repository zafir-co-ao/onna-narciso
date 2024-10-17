package scheduling

type DayliAppointmentsGetter interface {
	Get(day string) ([]Appointment, error)
}

type dayliAppointmentsGetterImpl struct {
	repo AppointmentRepository
}

func NewDayliAppointmentsGetter(repo AppointmentRepository) DayliAppointmentsGetter {
	return &dayliAppointmentsGetterImpl{repo: repo}
}

func (d *dayliAppointmentsGetterImpl) Get(date string) ([]Appointment, error) {
	return d.repo.FindByDate(date)
}
