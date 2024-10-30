package scheduling

import "github.com/zafir-co-ao/onna-narciso/internal/shared/id"

type AppointmentFinder interface {
	Execute(id string) (AppointmentOutput, error)
}

type appointmentFinderImpl struct {
	repo AppointmentRepository
}

func NewAppointmentFinder(r AppointmentRepository) AppointmentFinder {
	return &appointmentFinderImpl{repo: r}
}

func (f *appointmentFinderImpl) Execute(appointmentId string) (AppointmentOutput, error) {
	a, err := f.repo.FindByID(id.NewID(appointmentId))
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	return buildOutput(a), nil
}
