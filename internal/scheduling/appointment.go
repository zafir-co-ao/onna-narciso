package scheduling

import (
	"github.com/kindalus/godx/pkg/nanoid"
	_date "github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/hour"
)

const (
	StatusScheduled   Status = "scheduled"
	StatusCanceled    Status = "canceled"
	StatusRescheduled Status = "rescheduled"
)

var EmptyAppointment Appointment

type Service struct {
	ID   nanoid.ID
	Name Name
}

type Professional struct {
	ID          nanoid.ID
	Name        Name
	ServicesIDS []nanoid.ID
}

type Customer struct {
	ID          nanoid.ID
	Name        Name
	PhoneNumber string
}

type Status string

type Name string

func (n Name) String() string {
	return string(n)
}

type Appointment struct {
	ID               nanoid.ID
	ProfessionalID   nanoid.ID
	ProfessionalName Name
	CustomerID       nanoid.ID
	CustomerName     Name
	ServiceID        nanoid.ID
	ServiceName      Name
	Status           Status
	Date             _date.Date // Formato: 2024-10-01
	Hour             hour.Hour  // Formato 9:00
	Duration         int
}

func (a Appointment) GetID() nanoid.ID {
	return a.ID
}

func NewAppointment(
	id nanoid.ID,
	professionalID nanoid.ID,
	professionalName Name,
	customerID nanoid.ID,
	customerName Name,
	serviceID nanoid.ID,
	serviceName Name,
	date _date.Date,
	hour hour.Hour,
	duration int,
) (Appointment, error) {
	app := Appointment{
		ID:               id,
		ProfessionalID:   professionalID,
		ProfessionalName: professionalName,
		CustomerID:       customerID,
		CustomerName:     customerName,
		ServiceID:        serviceID,
		ServiceName:      serviceName,
		Date:             date,
		Hour:             hour,
		Duration:         duration,
		Status:           StatusScheduled,
	}

	return app, nil
}

func (a *Appointment) Reschedule(date string, time string, duration int) error {
	if !a.IsScheduled() {
		return ErrInvalidStatusToReschedule
	}

	d, err := _date.New(date)
	if err != nil {
		return err
	}

	h, err := hour.New(time)
	if err != nil {
		return err
	}

	a.Date = d
	a.Hour = h
	a.Duration = duration
	a.Status = StatusRescheduled

	return nil
}

func (a *Appointment) Cancel() error {
	if a.IsCancelled() {
		return ErrInvalidStatusToCancel
	}

	a.Status = StatusCanceled

	return nil
}

func (a *Appointment) IsScheduled() bool {
	return a.Status == StatusScheduled
}

func (a *Appointment) IsRescheduled() bool {
	return a.Status == StatusRescheduled
}

func (a *Appointment) IsCancelled() bool {
	return a.Status == StatusCanceled
}
