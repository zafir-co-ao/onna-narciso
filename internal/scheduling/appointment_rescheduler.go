package scheduling

type AppointmentRescheduler interface {
	Execute(id string, date string, hour string) (AppointmentOutput, error)
}

type appointmentRescheduler struct {
	repo AppointmentRepository
}

func NewAppointmentRescheduler(r AppointmentRepository) AppointmentRescheduler {
	return &appointmentRescheduler{repo: r}
}

func (r *appointmentRescheduler) Execute(id string, date string, hour string) (AppointmentOutput, error) {
	a, err := r.repo.FindByID(NewID(id))
	if err != nil {
		return AppointmentOutput{}, err
	}

	err = a.Reschedule(date, hour)
	if err != nil {
		return AppointmentOutput{}, err
	}

	r.repo.Save(a)

	return buildOutput(a), nil
}
