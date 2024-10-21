package scheduling

type AppointmentSchedulerDTO struct {
	ID             string
	ProfessionalID string
	CustomerID     string
	CustomerName   string
	CustomerPhone  string
	ServiceID      string
	Date           string
	StartHour      string
	Duration       int
}

type AppointmentScheduler interface {
	Schedule(d AppointmentSchedulerDTO) (string, error)
}

type appointmentScedulerImpl struct {
	customerRepo     CustomerRepository
	appointmentRepo  AppointmentRepository
	professionalRepo ProfessionalRepository
	serviceRepo      ServiceRepository
}

func NewAppointmentScheduler(aRepo AppointmentRepository, cRepo CustomerRepository, pRepo ProfessionalRepository, sRepo ServiceRepository) AppointmentScheduler {
	return &appointmentScedulerImpl{
		appointmentRepo:  aRepo,
		customerRepo:     cRepo,
		professionalRepo: pRepo,
		serviceRepo:      sRepo,
	}
}

func (s *appointmentScedulerImpl) Schedule(d AppointmentSchedulerDTO) (string, error) {
	_, err := s.professionalRepo.Get(d.ProfessionalID)
	if err != nil {
		return "", err
	}

	_, err = s.serviceRepo.Get(d.ServiceID)
	if err != nil {
		return "", err
	}

	customer, err := s.getOrAddCustomer(d)
	if err != nil {
		return "", err
	}

	app, _ := NewAppointmentBuilder().
		WithAppointmentID("1").
		WithProfessionalID(d.ProfessionalID).
		WithCustomerID(customer.ID).
		WithServiceID(d.ServiceID).
		WithDate(d.Date).
		WithStartHour(d.StartHour).
		WithDuration(d.Duration).
		Build()

	appointments, _ := s.appointmentRepo.FindByDate(d.Date)
	if !VerifyAvailability(app, appointments) {
		return "", ErrBusyTime
	}

	s.appointmentRepo.Save(app)

	return "1", nil
}

func (s *appointmentScedulerImpl) getOrAddCustomer(d AppointmentSchedulerDTO) (Customer, error) {
	if len(d.CustomerID) > 0 {
		c, err := s.customerRepo.Get(d.CustomerID)
		return c, err
	}

	customer := Customer{ID: "1000", Name: d.CustomerName, Phone: d.CustomerPhone}
	s.customerRepo.Save(customer)
	return customer, nil
}
