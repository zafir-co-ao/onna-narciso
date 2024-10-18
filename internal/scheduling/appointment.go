package scheduling

const (
	StatusScheduled Status = "scheduled"
)

type Status string

type Appointment struct {
	ID     string
	Status Status
	Date   string
	Start  string
}

func (a *Appointment) IsScheduled() bool {
	return a.Status == StatusScheduled
}
