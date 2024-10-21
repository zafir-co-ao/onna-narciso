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

	app, err := NewAppointmentBuilder().
		WithAppointmentID("1").
		WithProfessionalID(d.ProfessionalID).
		WithCustomerID(customer.ID).
		WithServiceID(d.ServiceID).
		WithDate(d.Date).
		WithStartHour(d.StartHour).
		WithDuration(d.Duration).
		Build()

	if err != nil {
		return "", err
	}

	appointments, _ := s.appointmentRepo.FindByDate(d.Date)
	if !VerifyAvailability(app, appointments) {
		return "", ErrBusyTime
	}

	s.appointmentRepo.Save(app)

	return "1", nil
}

func (s *appointmentScedulerImpl) getOrAddCustomer(d AppointmentSchedulerDTO) (Customer, error) {
	if len(d.CustomerID) > 0 {
		return s.customerRepo.Get(d.CustomerID)
	}

	if len(d.CustomerName) == 0 || len(d.CustomerPhone) == 0 {
		return EmptyCustomer, ErrCustomerRegistration
	}

	c := Customer{
		ID:    "1",
		Name:  d.CustomerName,
		Phone: d.CustomerPhone,
	}
	s.customerRepo.Save(c)
	return c, nil
}
