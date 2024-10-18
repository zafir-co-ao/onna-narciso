package scheduling

const (
	StatusScheduled Status = "scheduled"
)

type Status string

type Appointment struct {
	ID               string
	ProfessionalName string
	ProfessionalID   string
	CustomerID       string
	CustomerName     string
	ServiceName      string
	ServiceID        string
	Status           Status
	Date             string // Formato: 2024-10-01
	Start            string // Formato 9:00
	Duration         int
}

func (a *Appointment) IsScheduled() bool {
	return a.Status == StatusScheduled
}
