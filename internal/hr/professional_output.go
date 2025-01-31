package hr

type ServiceOutput struct {
	ID   string
	Name string
}

type ProfessionalOutput struct {
	ID       string
	Name     string
	Services []ServiceOutput
}

func toProfessionalOutput(p Professional) ProfessionalOutput {
	var services []ServiceOutput

	for _, s := range p.Services {
		r := ServiceOutput{
			ID:   s.ID.String(),
			Name: s.Name.String(),
		}
		services = append(services, r)
	}

	return ProfessionalOutput{
		ID:       p.ID.String(),
		Name:     p.Name.String(),
		Services: services,
	}
}
