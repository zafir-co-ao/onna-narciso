package scheduling

import "github.com/zafir-co-ao/onna-narciso/internal/shared/id"

type AppointmentGetter interface {
	Get(id string) (AppointmentOutput, error)
}

type appointmentGetterImpl struct {
	repo AppointmentRepository
}

func NewAppointmentGetter(r AppointmentRepository) AppointmentGetter {
	return &appointmentGetterImpl{repo: r}
}

func (u *appointmentGetterImpl) Get(appointmentId string) (AppointmentOutput, error) {
	a, err := u.repo.FindByID(id.NewID(appointmentId))
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	return toAppointmentOutput(a), nil
}
