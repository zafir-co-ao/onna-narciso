package scheduling

type Appointment struct {
	ID               string
	ServiceID        string
	ServiceName      string
	ProfessionalID   string
	ProfessionalName string
	Status           string
	Date             string // Formato: 2024-10-01
	Start            string // Formato 9:00
	Duration         int
	CustomerID       string
	CustomerName     string
}
