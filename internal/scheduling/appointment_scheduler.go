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
	Schedule(d AppointmentSchedulerInput) (AppointmentOutput, error)
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

func (u *appointmentScedulerImpl) Schedule(d AppointmentSchedulerInput) (AppointmentOutput, error) {
	p, err := u.professionalAcl.FindProfessionalByID(d.ProfessionalID)
	if err != nil {
		return AppointmentOutput{}, err
	}

	s, err := u.serviceAcl.FindServiceByID(d.ServiceID)
	if err != nil {
		return AppointmentOutput{}, err
	}

	c, err := u.findOrRegistrationCustomer(d)
	if err != nil {
		return AppointmentOutput{}, err
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
		return AppointmentOutput{}, err
	}

	appointments, _ := u.repo.FindByDateAndStatus(d.Date, StatusScheduled)
	if !VerifyAvailability(app, appointments) {
		return AppointmentOutput{}, ErrBusyTime
	}

	u.repo.Save(app)

	return buildOutput(app), nil
}

func (u *appointmentScedulerImpl) findOrRegistrationCustomer(d AppointmentSchedulerInput) (Customer, error) {
	if len(d.CustomerID) > 0 {
		return u.customerAcl.FindCustomerByID(d.CustomerID)
	}

	return u.customerAcl.RequestCustomerRegistration(d.CustomerName, d.CustomerPhone)
}
