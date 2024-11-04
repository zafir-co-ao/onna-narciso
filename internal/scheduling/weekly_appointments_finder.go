package scheduling

import (
	"github.com/kindalus/gofunc/pkg/collections"
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

func (w *weeklyAppointmentsGetterImpl) Find(week string, serviceID string, professionalsIDs []string) ([]AppointmentOutput, error) {
	a, e := w.repo.FindByWeekServiceAndProfessionals(week, serviceID, professionalsIDs)
	if e != nil {
		return nil, e
	}

	return collections.Map(a, toAppointmentOutput), nil
}
