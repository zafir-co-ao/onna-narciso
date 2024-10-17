package scheduling


type DateIsSpecificantion struct {
	Date string
}

func (d DateIsSpecificantion) IsSatisfiedBy(a Appointment) bool {
	return a.Date == d.Date
}

