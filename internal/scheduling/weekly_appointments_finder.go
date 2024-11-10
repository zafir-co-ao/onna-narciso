package scheduling

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
)

type WeeklyAppointmentsFinder interface {
	Find(week string, serviceID string, professionalsIDs []string) ([]AppointmentOutput, error)
}

type weeklyAppointmentsGetterImpl struct {
	repo AppointmentRepository
}

func NewWeeklyAppointmentsGetter(repo AppointmentRepository) WeeklyAppointmentsFinder {
	return &weeklyAppointmentsGetterImpl{repo: repo}
}

func (w *weeklyAppointmentsGetterImpl) Find(adate string, serviceID string, professionalsIDs []string) ([]AppointmentOutput, error) {

	pids := xslices.Map(professionalsIDs, func(x string) nanoid.ID { return nanoid.ID(x) })

	a, e := w.repo.FindByWeekServiceAndProfessionals(date.Date(adate), nanoid.ID(serviceID), pids)
	if e != nil {
		return nil, e
	}

	return xslices.Map(a, toAppointmentOutput), nil
}
