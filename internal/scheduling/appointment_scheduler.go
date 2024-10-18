package scheduling

type AppointmentScheduler interface {
	Schedule(professionalId, customerId, serviceId string) (string, error)
}

type appointmentScedulerImpl struct {
	repo AppointmentRepository
}

func NewAppointmentScheduler(repo AppointmentRepository) AppointmentScheduler {
	return &appointmentScedulerImpl{
		repo: repo,
	}
}

func (s *appointmentScedulerImpl) Schedule(professionalId, customerId, serviceId string) (string, error) {

	app := Appointment{
		ID:             "1",
		Status:         StatusScheduled,
		ProfessionalID: professionalId,
		CustomerID:     customerId,
		ServiceID:      serviceId,
	}

	s.repo.Save(app)

	return "1", nil
}
