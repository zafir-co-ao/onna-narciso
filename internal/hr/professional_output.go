package hr

type ProfessionalOutput struct {
	ID   string
	Name string
}

func toProfessionalOutput(p Professional) ProfessionalOutput {
	return ProfessionalOutput{
		ID:   p.ID.String(),
		Name: p.Name.String(),
	}
}
