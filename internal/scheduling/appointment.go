package scheduling

import (
	"strconv"
	"strings"

	"github.com/zafir-co-ao/onna-narciso/internal/shared/id"
)

const (
	StatusScheduled   Status = "scheduled"
	StatusCanceled    Status = "canceled"
	StatusRescheduled Status = "rescheduled"
)

var EmptyAppointment = Appointment{}

type Service struct {
	ID   id.ID
	Name Name
}

type Professional struct {
	ID          id.ID
	Name        Name
	ServicesIDS []id.ID
}

type Customer struct {
	ID          id.ID
	Name        Name
	PhoneNumber string
}

type Status string
type Name string

func (n Name) String() string {
	return string(n)
}

type Appointment struct {
	ID               id.ID
	ProfessionalID   id.ID
	ProfessionalName Name
	CustomerID       id.ID
	CustomerName     Name
	ServiceID        id.ID
	ServiceName      Name
	Status           Status
	Date             Date // Formato: 2024-10-01
	Hour             Hour // Formato 9:00
	End              Hour
	Duration         int
}

func NewAppointment(
	ID id.ID,
	ProfessionalID id.ID,
	ProfessionalName Name,
	CustomerID id.ID,
	CustomerName Name,
	ServiceID id.ID,
	ServiceName Name,
	Date Date,
	Start Hour,
	Duration int,
) (Appointment, error) {
	app := Appointment{
		ID:               ID,
		ProfessionalID:   ProfessionalID,
		ProfessionalName: ProfessionalName,
		CustomerID:       CustomerID,
		CustomerName:     CustomerName,
		ServiceID:        ServiceID,
		ServiceName:      ServiceName,
		Date:             Date,
		Hour:             Start,
		Duration:         Duration,
		Status:           StatusScheduled,
	}

	app.calculateEnd()

	return app, nil
}

func (a *Appointment) Reschedule(date string, hour string, duration int) error {
	if !a.IsScheduled() {
		return ErrInvalidStatusToReschedule
	}

	d, err := NewDate(date)
	if err != nil {
		return err
	}

	h, err := NewHour(hour)
	if err != nil {
		return err
	}

	a.Date = d
	a.Hour = h
	a.Duration = duration
	a.Status = StatusRescheduled

	a.calculateEnd()
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

func (a *Appointment) calculateEnd() {
	parts := strings.Split(a.Hour.Value(), ":")
	hour, _ := strconv.ParseInt(parts[0], 10, 8)
	minutes, _ := strconv.ParseInt(parts[1], 10, 8)

	totalMinutes := a.Duration + int(minutes)
	endHour := hour + int64(totalMinutes)/60
	endMinutes := totalMinutes % 60

	if endMinutes < 10 {
		a.End, _ = NewHour(strconv.Itoa(int(endHour)) + ":0" + strconv.Itoa(int(endMinutes)))
		return
	}

	a.End, _ = NewHour(strconv.Itoa(int(endHour)) + ":" + strconv.Itoa(int(endMinutes)))
}
