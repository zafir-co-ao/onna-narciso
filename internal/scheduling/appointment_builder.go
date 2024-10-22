package scheduling

type AppointmentBuilder interface {
	WithAppointmentID(id ID) AppointmentBuilder
	WithProfessionalID(id string) AppointmentBuilder
	WithCustomerID(id string) AppointmentBuilder
	WithServiceID(id string) AppointmentBuilder
	WithDate(date string) AppointmentBuilder
	WithStartHour(hour string) AppointmentBuilder
	WithDuration(duration int) AppointmentBuilder
	Build() (Appointment, error)
}

type appointmentBuilder struct {
	ID             ID
	ProfessionalID string
	CustomerID     string
	ServiceID      string
	Date           string
	StartHour      string
	Duration       int
}

func NewAppointmentBuilder() AppointmentBuilder {
	return &appointmentBuilder{}
}

func (b *appointmentBuilder) WithAppointmentID(id ID) AppointmentBuilder {
	b.ID = id
	return b
}

func (b *appointmentBuilder) WithProfessionalID(id string) AppointmentBuilder {
	b.ProfessionalID = id
	return b
}

func (b *appointmentBuilder) WithCustomerID(id string) AppointmentBuilder {
	b.CustomerID = id
	return b
}

func (b *appointmentBuilder) WithServiceID(id string) AppointmentBuilder {
	b.ServiceID = id
	return b
}

func (b *appointmentBuilder) WithDate(date string) AppointmentBuilder {
	b.Date = date
	return b
}

func (b *appointmentBuilder) WithStartHour(hour string) AppointmentBuilder {
	b.StartHour = hour
	return b
}

func (b *appointmentBuilder) WithDuration(duration int) AppointmentBuilder {
	b.Duration = duration
	return b
}

func (b *appointmentBuilder) Build() (Appointment, error) {
	return NewAppointment(
		b.ID,
		b.ProfessionalID,
		b.CustomerID,
		b.ServiceID,
		b.Date,
		b.StartHour,
		b.Duration,
	)
}
