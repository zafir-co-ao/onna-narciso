package scheduling

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/hour"
)

type AppointmentBuilder interface {
	WithAppointmentID(id nanoid.ID) AppointmentBuilder
	WithProfessionalID(id nanoid.ID) AppointmentBuilder
	WithProfessionalName(name Name) AppointmentBuilder
	WithCustomerID(id nanoid.ID) AppointmentBuilder
	WithCustomerName(name Name) AppointmentBuilder
	WithServiceID(id nanoid.ID) AppointmentBuilder
	WithServiceName(name Name) AppointmentBuilder
	WithDate(date string) AppointmentBuilder
	WithHour(hour string) AppointmentBuilder
	WithDuration(duration int) AppointmentBuilder
	Build() (Appointment, error)
}

type appointmentBuilder struct {
	ID               nanoid.ID
	ProfessionalID   nanoid.ID
	ProfessionalName Name
	CustomerID       nanoid.ID
	CustomerName     Name
	ServiceID        nanoid.ID
	ServiceName      Name
	Date             string
	Hour             string
	Duration         int
}

func NewAppointmentBuilder() AppointmentBuilder {
	return &appointmentBuilder{}
}

func (b *appointmentBuilder) WithAppointmentID(id nanoid.ID) AppointmentBuilder {
	b.ID = id
	return b
}

func (b *appointmentBuilder) WithProfessionalID(id nanoid.ID) AppointmentBuilder {
	b.ProfessionalID = id
	return b
}

func (b *appointmentBuilder) WithProfessionalName(name Name) AppointmentBuilder {
	b.ProfessionalName = name
	return b
}

func (b *appointmentBuilder) WithCustomerID(id nanoid.ID) AppointmentBuilder {
	b.CustomerID = id
	return b
}

func (b *appointmentBuilder) WithCustomerName(name Name) AppointmentBuilder {
	b.CustomerName = name
	return b
}

func (b *appointmentBuilder) WithServiceID(id nanoid.ID) AppointmentBuilder {
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
	date, err := date.New(b.Date)
	if err != nil {
		return EmptyAppointment, err
	}

	hour, err := hour.New(b.Hour)
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
