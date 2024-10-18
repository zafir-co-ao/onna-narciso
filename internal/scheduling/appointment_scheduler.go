package scheduling

type AppointmentScheduler interface {
	Schedule(professionalId string) (string, error)
}

type appointmentScedulerImpl struct {
	repo AppointmentRepository
}

func NewAppointmentScheduler(repo AppointmentRepository) AppointmentScheduler {
	return &appointmentScedulerImpl{
		repo: repo,
	}
}

func (s *appointmentScedulerImpl) Schedule(professionalId string) (string, error) {

	app := Appointment{
		ID:             "1",
		Status:         StatusScheduled,
		ProfessionalID: professionalId,
	}

	s.repo.Save(app)

	return "1", nil
}
