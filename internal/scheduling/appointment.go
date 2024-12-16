package scheduling

import (
	"slices"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/duration"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/hour"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

const (
	StatusScheduled   Status = "Agendado"
	StatusRescheduled Status = "Reagendado"
	StatusCanceled    Status = "Cancelado"
	StatusClosed      Status = "Fechado"
)

var EmptyAppointment = Appointment{}

type Service struct {
	ID       nanoid.ID
	Name     name.Name
	Duration duration.Duration
}

type Professional struct {
	ID          nanoid.ID
	Name        name.Name
	ServicesIDS []nanoid.ID
}

func (p Professional) ContainsService(sid nanoid.ID) bool {
	return slices.Contains(p.ServicesIDS, sid)
}

type Customer struct {
	ID          nanoid.ID `json:"id"`
	Name        name.Name `json:"name"`
	PhoneNumber string    `json:"phoneNumber"`
	Nif         string    `json:"nif"`
}

type Status string

type Appointment struct {
	ID               nanoid.ID
	ProfessionalID   nanoid.ID
	ProfessionalName name.Name
	CustomerID       nanoid.ID
	CustomerName     name.Name
	ServiceID        nanoid.ID
	ServiceName      name.Name
	Status           Status
	Date             date.Date // Formato: 2024-10-01
	Hour             hour.Hour // Formato 9:00
	Duration         duration.Duration
}

func (a Appointment) GetID() nanoid.ID {
	return a.ID
}

func (a *Appointment) Reschedule() error {
	if !a.IsScheduled() && !a.IsRescheduled() {
		return ErrInvalidStatusToReschedule
	}

	a.Status = StatusRescheduled

	return nil
}

func (a *Appointment) Close() error {
	if a.IsClosed() {
		return ErrInvalidStatusToClose
	}

	a.Status = StatusClosed

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

func (a *Appointment) IsClosed() bool {
	return a.Status == StatusClosed
}

func (a *Appointment) IsCancelled() bool {
	return a.Status == StatusCanceled
}
