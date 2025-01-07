package crm

type CustomerOutput struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Nif         string `json:"nif"`
	BirthDate   string
	Email       string
	PhoneNumber string `json:"phoneNumber"`
}

func toCustomerOutput(c Customer) CustomerOutput {
	return CustomerOutput{
		ID:          c.ID.String(),
		Name:        c.Name.String(),
		Nif:         c.Nif.String(),
		BirthDate:   c.BirthDate.String(),
		Email:       c.Email.String(),
		PhoneNumber: c.PhoneNumber.String(),
	}
}
