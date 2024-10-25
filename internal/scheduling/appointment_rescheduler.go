package scheduling

type AppointmentRescheduler interface {
	Execute(id string) (AppointmentOutput, error)
}

type appointmentRescheduler struct {
	repo AppointmentRepository
}

func NewAppointmentRescheduler(r AppointmentRepository) AppointmentRescheduler {
	return &appointmentRescheduler{repo: r}
}

func (r *appointmentRescheduler) Execute(id string) (AppointmentOutput, error) {
	a, _ := r.repo.FindByID(NewID(id))

	a.Reschedule()

	r.repo.Save(a)

	return buildOutput(a), nil
}
