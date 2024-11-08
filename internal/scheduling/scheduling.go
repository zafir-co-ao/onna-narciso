package scheduling

import "errors"

var (
	ErrBusyTime                  = errors.New("Schedule time is busy")
	ErrInvalidStatusToCancel     = errors.New("Invalid status to cancel")
	ErrInvalidStatusToReschedule = errors.New("Invalid status to reschedule")
)

var EmptyAppointmentOutput = AppointmentOutput{}

type AppointmentOutput struct {
	ID               string
	CustomerID       string
	CustomerName     string
	ServiceID        string
	ServiceName      string
	ProfessionalID   string
	ProfessionalName string
	Date             string
	Hour             string
	Duration         int
}

func toAppointmentOutput(a Appointment) AppointmentOutput {
	return AppointmentOutput{
		ID:               a.ID.String(),
		CustomerID:       string(a.CustomerID),
		CustomerName:     string(a.CustomerName),
		ServiceID:        string(a.ServiceID),
		ServiceName:      string(a.ServiceName),
		ProfessionalID:   string(a.ProfessionalID),
		ProfessionalName: string(a.ProfessionalName),
		Date:             a.Date.Value(),
		Hour:             a.Hour.Value(),
		Duration:         a.Duration,
	}
}
