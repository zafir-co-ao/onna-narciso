package scheduling

import "github.com/zafir-co-ao/onna-narciso/internal/shared/id"

type AppointmentBuilder interface {
	WithAppointmentID(id id.ID) AppointmentBuilder
	WithProfessionalID(id id.ID) AppointmentBuilder
	WithProfessionalName(name Name) AppointmentBuilder
	WithCustomerID(id id.ID) AppointmentBuilder
	WithCustomerName(name Name) AppointmentBuilder
	WithServiceID(id id.ID) AppointmentBuilder
	WithServiceName(name Name) AppointmentBuilder
	WithDate(date string) AppointmentBuilder
	WithHour(hour string) AppointmentBuilder
	WithDuration(duration int) AppointmentBuilder
	Build() (Appointment, error)
}

type appointmentBuilder struct {
	ID               id.ID
	ProfessionalID   id.ID
	ProfessionalName Name
	CustomerID       id.ID
	CustomerName     Name
	ServiceID        id.ID
	ServiceName      Name
	Date             string
	Hour             string
	Duration         int
}

func NewAppointmentBuilder() AppointmentBuilder {
	return &appointmentBuilder{}
}

func (b *appointmentBuilder) WithAppointmentID(id id.ID) AppointmentBuilder {
	b.ID = id
	return b
}

func (b *appointmentBuilder) WithProfessionalID(id id.ID) AppointmentBuilder {
	b.ProfessionalID = id
	return b
}

func (b *appointmentBuilder) WithProfessionalName(name Name) AppointmentBuilder {
	b.ProfessionalName = name
	return b
}

func (b *appointmentBuilder) WithCustomerID(id id.ID) AppointmentBuilder {
	b.CustomerID = id
	return b
}

func (b *appointmentBuilder) WithCustomerName(name Name) AppointmentBuilder {
	b.CustomerName = name
	return b
}

func (b *appointmentBuilder) WithServiceID(id id.ID) AppointmentBuilder {
	b.ServiceID = id
	return b
}

func (b *appointmentBuilder) WithServiceName(name Name) AppointmentBuilder {
	b.ServiceName = name
	return b
}

func (b *appointmentBuilder) WithDate(date string) AppointmentBuilder {
	b.Date = date
	return b
}

func (b *appointmentBuilder) WithHour(hour string) AppointmentBuilder {
	b.Hour = hour
	return b
}

func (b *appointmentBuilder) WithDuration(duration int) AppointmentBuilder {
	b.Duration = duration
	return b
}

func (b *appointmentBuilder) Build() (Appointment, error) {
	date, err := NewDate(b.Date)
	if err != nil {
		return EmptyAppointment, err
	}

	hour, err := NewHour(b.Hour)
	if err != nil {
		return EmptyAppointment, err
	}

	return NewAppointment(
		b.ID,
		b.ProfessionalID,
		b.ProfessionalName,
		b.CustomerID,
		b.CustomerName,
		b.ServiceID,
		b.ServiceName,
		date,
		hour,
		b.Duration,
	)
}
