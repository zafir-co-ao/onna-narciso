package scheduling

type AppointmentRepository interface {
	Get(id string) (Appointment, error)
	Save(appointment Appointment) error
	FindByDate(date string) ([]Appointment, error)
}
