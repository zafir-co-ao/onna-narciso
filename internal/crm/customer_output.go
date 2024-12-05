package crm

type CustomerOutput struct {
	ID          string
	Name        string
	Nif         string
	BirthDate   string
	Email       string
	PhoneNumber string
}

func toCustomerOutput(c Customer) CustomerOutput {
	return CustomerOutput{
		ID:          c.ID.String(),
		Name:        c.Name.String(),
		Nif:         c.Nif.String(),
		BirthDate:   c.BirthDate.String(),
		Email:       c.Email,
		PhoneNumber: c.PhoneNumber,
	}
}
