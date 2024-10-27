package scheduling

type AppointmentReschedulerInput struct {
	ID        string
	Date      string
	StartHour string
	Duration  int
}

type AppointmentRescheduler interface {
	Execute(i AppointmentReschedulerInput) (AppointmentOutput, error)
}

type appointmentRescheduler struct {
	repo AppointmentRepository
}

func NewAppointmentRescheduler(r AppointmentRepository) AppointmentRescheduler {
	return &appointmentRescheduler{repo: r}
}

func (r *appointmentRescheduler) Execute(i AppointmentReschedulerInput) (AppointmentOutput, error) {
	a, err := r.repo.FindByID(NewID(i.ID))
	if err != nil {
		return AppointmentOutput{}, err
	}

	err = a.Reschedule(i.Date, i.StartHour, i.Duration)
	if err != nil {
		return AppointmentOutput{}, err
	}

	appointments, err := r.repo.FindByDateAndStatus(i.Date, StatusScheduled)
	if err != nil {
		return AppointmentOutput{}, err
	}

	if !VerifyAvailability(a, appointments) {
		return AppointmentOutput{}, ErrBusyTime
	}

	r.repo.Save(a)

	return buildOutput(a), nil
}
