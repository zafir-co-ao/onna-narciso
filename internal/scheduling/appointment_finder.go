package scheduling

import "github.com/kindalus/godx/pkg/nanoid"

type AppointmentFinder interface {
	FindByID(id string) (AppointmentOutput, error)
}

type appointmentFinderImpl struct {
	repo AppointmentRepository
}

func NewAppointmentFinder(r AppointmentRepository) AppointmentFinder {
	return &appointmentFinderImpl{repo: r}
}

func (f *appointmentFinderImpl) FindByID(appointmentId string) (AppointmentOutput, error) {
	a, err := f.repo.FindByID(nanoid.ID(appointmentId))
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	return toAppointmentOutput(a), nil
}
