package scheduling

type ProfessionalRepository interface {
	Get(id string) (Professional, error)
	Save(p Professional) error
}
