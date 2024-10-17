package scheduling

type AppointmentScheduler interface {
	Schedule() (string, error)
}

type appointmentScedulerImpl struct {
	repo AppointmentRepository
}

func NewAppointmentScheduler(repo AppointmentRepository) AppointmentScheduler {
	return &appointmentScedulerImpl{
		repo: repo,
	}
}

func (s *appointmentScedulerImpl) Schedule() (string, error) {

	app := Appointment{
		ID:     "1",
		Status: "Scheduled",
	}

	s.repo.Save(app)

	return "1", nil
}
