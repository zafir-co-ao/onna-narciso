package scheduling

import "errors"

var (
	ErrBusyTime                  = errors.New("Schedule time is busy")
	ErrInvalidStatusToCancel     = errors.New("Invalid status to cancel")
	ErrInvalidStatusToReschedule = errors.New("Invalid status to reschedule")
)

type AppointmentOutput struct {
	ID               string
	CustomerName     string
	ServiceName      string
	ProfessionalName string
	Date             string
	Hour             string
	Start            int
	Duration         int
}

func buildOutput(a Appointment) AppointmentOutput {
	o := AppointmentOutput{
		ID:               a.ID.Value(),
		CustomerName:     a.CustomerName,
		ServiceName:      a.ServiceName,
		ProfessionalName: a.ProfessionalName,
		Date:             a.Date.Value(),
		Hour:             a.Start.Value(),
		Duration:         a.Duration,
	}
	return o
}
