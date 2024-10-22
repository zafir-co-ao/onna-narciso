package scheduling

type AppointmentBuilder interface {
	WithAppointmentID(id ID) AppointmentBuilder
	WithProfessionalID(id ID) AppointmentBuilder
	WithCustomerID(id ID) AppointmentBuilder
	WithServiceID(id ID) AppointmentBuilder
	WithDate(date string) AppointmentBuilder
	WithStartHour(hour string) AppointmentBuilder
	WithDuration(duration int) AppointmentBuilder
	Build() (Appointment, error)
}

type appointmentBuilder struct {
	ID             ID
	ProfessionalID ID
	CustomerID     ID
	ServiceID      ID
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

func (b *appointmentBuilder) WithProfessionalID(id ID) AppointmentBuilder {
	b.ProfessionalID = id
	return b
}

func (b *appointmentBuilder) WithCustomerID(id ID) AppointmentBuilder {
	b.CustomerID = id
	return b
}

func (b *appointmentBuilder) WithServiceID(id ID) AppointmentBuilder {
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
	date, err := NewDate(b.Date)
	if err != nil {
		return Appointment{}, err
	}

	hour, err := NewHour(b.StartHour)
	if err != nil {
		return Appointment{}, err
	}

	return NewAppointment(
		b.ID,
		b.ProfessionalID,
		b.CustomerID,
		b.ServiceID,
		date,
		hour,
		b.Duration,
	)
}
