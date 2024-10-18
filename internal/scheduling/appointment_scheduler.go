package scheduling

type AppointmentSchedulerDTO struct {
	ID             string
	ProfessionalID string
	CustomerID     string
	ServiceID      string
	Date           string
	StartHour      string
	Duration       int
}

type AppointmentScheduler interface {
	Schedule(d AppointmentSchedulerDTO) (string, error)
}

type appointmentScedulerImpl struct {
	repo AppointmentRepository
}

func NewAppointmentScheduler(repo AppointmentRepository) AppointmentScheduler {
	return &appointmentScedulerImpl{
		repo: repo,
	}
}

func (s *appointmentScedulerImpl) Schedule(d AppointmentSchedulerDTO) (string, error) {

	app := Appointment{
		ID:             "1",
		Status:         StatusScheduled,
		ProfessionalID: d.ProfessionalID,
		CustomerID:     d.CustomerID,
		ServiceID:      d.ServiceID,
		Date:           d.Date,
		Start:          d.StartHour,
		Duration:       d.Duration,
	}

	s.repo.Save(app)

	return "1", nil
}
