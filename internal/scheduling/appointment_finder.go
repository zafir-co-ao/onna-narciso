package scheduling

import "github.com/kindalus/godx/pkg/nanoid"

type AppointmentFinder interface {
	FindByID(id string) (AppointmentOutput, error)
}

type appointmentFinderImpl struct {
	repo AppointmentRepository
}

func NewAppointmentFinder(repo AppointmentRepository) AppointmentFinder {
	return &appointmentFinderImpl{repo}
}

func (f *appointmentFinderImpl) FindByID(id string) (AppointmentOutput, error) {
	a, err := f.repo.FindByID(nanoid.ID(id))
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	return toAppointmentOutput(a), nil
}
