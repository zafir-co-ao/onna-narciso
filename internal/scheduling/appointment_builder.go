package scheduling

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/hour"
)

type appointmentBuilder Appointment

func NewAppointmentBuilder() *appointmentBuilder {
	return &appointmentBuilder{
		ID:     nanoid.New(),
		Status: StatusScheduled,
	}
}

func (b *appointmentBuilder) WithAppointmentID(id nanoid.ID) *appointmentBuilder {
	b.ID = id
	return b
}

func (b *appointmentBuilder) WithProfessional(id nanoid.ID, name Name) *appointmentBuilder {
	b.ProfessionalID = id
	b.ProfessionalName = name
	return b
}

func (b *appointmentBuilder) WithCustomer(id nanoid.ID, name Name) *appointmentBuilder {
	b.CustomerID = id
	b.CustomerName = name
	return b
}

func (b *appointmentBuilder) WithService(id nanoid.ID, name Name) *appointmentBuilder {
	b.ServiceID = id
	b.ServiceName = name
	return b
}

func (b *appointmentBuilder) WithDate(date date.Date) *appointmentBuilder {
	b.Date = date
	return b
}

func (b *appointmentBuilder) WithStatus(status Status) *appointmentBuilder {
	b.Status = status
	return b
}

func (b *appointmentBuilder) WithHour(hour hour.Hour) *appointmentBuilder {
	b.Hour = hour
	return b
}

func (b *appointmentBuilder) WithDuration(duration int) *appointmentBuilder {
	b.Duration = duration
	return b
}

func (b appointmentBuilder) Build() Appointment {
	return Appointment(b)
}
