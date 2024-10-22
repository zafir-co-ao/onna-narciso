package scheduling

type AppointmentSchedulerInput struct {
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
	Schedule(d AppointmentSchedulerInput) (string, error)
}

type appointmentScedulerImpl struct {
	repo            AppointmentRepository
	serviceAcl      ServiceAcl
	customerAcl     CustomerAcl
	professionalAcl ProfessionalAcl
}

func NewAppointmentScheduler(repo AppointmentRepository, cacl CustomerAcl, pacl ProfessionalAcl, sacl ServiceAcl) AppointmentScheduler {
	return &appointmentScedulerImpl{
		repo:            repo,
		customerAcl:     cacl,
		professionalAcl: pacl,
		serviceAcl:      sacl,
	}
}

func (u *appointmentScedulerImpl) Schedule(d AppointmentSchedulerInput) (string, error) {
	p, err := u.professionalAcl.FindProfessionalByID(d.ProfessionalID)
	if err != nil {
		return "", err
	}

	s, err := u.serviceAcl.FindServiceByID(d.ServiceID)
	if err != nil {
		return "", err
	}

	c, err := u.getOrAddCustomer(d)
	if err != nil {
		return "", err
	}

	id, _ := Random()

	app, err := NewAppointmentBuilder().
		WithAppointmentID(id).
		WithProfessionalID(NewID(p.ID)).
		WithProfessionalName(p.Name).
		WithCustomerID(NewID(c.ID)).
		WithCustomerName(c.Name).
		WithServiceID(NewID(s.ID)).
		WithServiceName(s.Name).
		WithDate(d.Date).
		WithStartHour(d.StartHour).
		WithDuration(d.Duration).
		Build()

	if err != nil {
		return "", err
	}

	appointments, _ := u.repo.FindByDate(d.Date)
	if !VerifyAvailability(app, appointments) {
		return "", ErrBusyTime
	}

	u.repo.Save(app)

	return id.Value(), nil
}

func (u *appointmentScedulerImpl) getOrAddCustomer(d AppointmentSchedulerInput) (Customer, error) {
	if len(d.CustomerID) > 0 {
		return u.customerAcl.FindCustomerByID(d.CustomerID)
	}

	if len(d.CustomerName) == 0 || len(d.CustomerPhone) == 0 {
		return Customer{}, ErrCustomerRegistration
	}

	c := Customer{
		ID:   "1",
		Name: d.CustomerName,
	}
	return c, nil
}
