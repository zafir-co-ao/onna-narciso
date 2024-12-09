package services

type ServiceOutput struct {
	ID          string
	Name        string
	Description string
	Price       string
	Duration    int
}

func toServiceOutput(s Service) ServiceOutput {
	return ServiceOutput{
		ID:          s.ID.String(),
		Name:        s.Name.String(),
		Duration:    s.Duration.Value(),
		Description: string(s.Description),
		Price:       string(s.Price),
	}
}
