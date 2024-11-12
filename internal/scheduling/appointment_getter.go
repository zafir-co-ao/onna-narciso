package scheduling

import "github.com/kindalus/godx/pkg/nanoid"

type AppointmentGetter interface {
	Get(id string) (AppointmentOutput, error)
}

type appointmentGetterImpl struct {
	repo AppointmentRepository
}

func NewAppointmentGetter(r AppointmentRepository) AppointmentGetter {
	return &appointmentGetterImpl{repo: r}
}

func (f *appointmentGetterImpl) Get(appointmentId string) (AppointmentOutput, error) {
	a, err := f.repo.FindByID(nanoid.ID(appointmentId))
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	return toAppointmentOutput(a), nil
}
