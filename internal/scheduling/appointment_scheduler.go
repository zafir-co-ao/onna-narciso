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

type AppointmentSchedulerOutput struct {
	ID               string
	CustomerName     string
	ServiceName      string
	ProfessionalName string
	Date             string
	Hour             string
	Start            int
	Duration         int
}

type AppointmentScheduler interface {
	Schedule(d AppointmentSchedulerInput) (AppointmentSchedulerOutput, error)
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

func (u *appointmentScedulerImpl) Schedule(d AppointmentSchedulerInput) (AppointmentSchedulerOutput, error) {
	p, err := u.professionalAcl.FindProfessionalByID(d.ProfessionalID)
	if err != nil {
		return AppointmentSchedulerOutput{}, err
	}

	s, err := u.serviceAcl.FindServiceByID(d.ServiceID)
	if err != nil {
		return AppointmentSchedulerOutput{}, err
	}

	c, err := u.findOrRegistrationCustomer(d)
	if err != nil {
		return AppointmentSchedulerOutput{}, err
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
		return AppointmentSchedulerOutput{}, err
	}

	appointments, _ := u.repo.FindByDateAndStatus(d.Date, StatusScheduled)
	if !VerifyAvailability(app, appointments) {
		return AppointmentSchedulerOutput{}, ErrBusyTime
	}

	u.repo.Save(app)

	return u.buildOutput(app), nil
}

func (u *appointmentScedulerImpl) findOrRegistrationCustomer(d AppointmentSchedulerInput) (Customer, error) {
	if len(d.CustomerID) > 0 {
		return u.customerAcl.FindCustomerByID(d.CustomerID)
	}

	return u.customerAcl.RequestCustomerRegistration(d.CustomerName, d.CustomerPhone)
}

func (*appointmentScedulerImpl) buildOutput(a Appointment) AppointmentSchedulerOutput {
	o := AppointmentSchedulerOutput{
		ID:               a.ID.Value(),
		CustomerName:     a.CustomerName,
		ServiceName:      a.ServiceName,
		ProfessionalName: a.ProfessionalName,
		Date:             a.Date.Value(),
		Hour:             a.Start.Value(),
		Duration:         a.Duration,
	}
	return o
}
