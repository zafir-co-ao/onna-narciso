package scheduling

import (
	"strconv"
	"strings"
)

const (
	StatusScheduled Status = "scheduled"
	StatusCancelled Status = "cancelled"
)

type Service struct {
	ID   string
	Name string
}

type Professional struct {
	ID   string
	Name string
}

type Customer struct {
	ID   string
	Name string
}

type Status string

type Appointment struct {
	ID               ID
	ProfessionalID   ID
	ProfessionalName string
	CustomerID       ID
	CustomerName     string
	ServiceID        ID
	ServiceName      string
	Status           Status
	Date             Date // Formato: 2024-10-01
	Start            Hour // Formato 9:00
	End              Hour
	Duration         int
}

func NewAppointment(
	ID ID,
	ProfessionalID ID,
	ProfessionalName string,
	CustomerID ID,
	ServiceID ID,
	Date Date,
	Start Hour,
	Duration int,
) (Appointment, error) {
	app := Appointment{
		ID:               ID,
		ProfessionalID:   ProfessionalID,
		ProfessionalName: ProfessionalName,
		CustomerID:       CustomerID,
		ServiceID:        ServiceID,
		Date:             Date,
		Start:            Start,
		Duration:         Duration,
		Status:           StatusScheduled,
	}

	app.calculateEnd()

	return app, nil
}

func (a *Appointment) Cancel() {
	a.Status = StatusCancelled
}

func (a *Appointment) IsScheduled() bool {
	return a.Status == StatusScheduled
}

func (a *Appointment) IsCancelled() bool {
	return a.Status == StatusCancelled
}

func (a *Appointment) calculateEnd() {
	parts := strings.Split(a.Start.Value(), ":")
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
