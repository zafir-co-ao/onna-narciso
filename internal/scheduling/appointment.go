package scheduling

const (
	StatusScheduled Status = "scheduled"
)

type Status string

type Appointment struct {
	ID             string
	ProfessionalID string
	CustomerID     string
	ServiceID      string
	Status         Status
	Date           string
	Start          string
	Duration       int
}

func (a *Appointment) IsScheduled() bool {
	return a.Status == StatusScheduled
}
