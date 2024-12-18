package scheduling

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/kindalus/godx/pkg/xslices"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
)

type WeeklyAppointmentsFinder interface {
	Find(week string, serviceID string, professionalsIDs []string) ([]AppointmentOutput, error)
}

type weeklyAppointmentsFinderImpl struct {
	repo AppointmentRepository
}

func NewWeeklyAppointmentsFinder(repo AppointmentRepository) WeeklyAppointmentsFinder {
	return &weeklyAppointmentsFinderImpl{repo}
}

func (w *weeklyAppointmentsFinderImpl) Find(adate string, serviceID string, professionalsIDs []string) ([]AppointmentOutput, error) {

	pids := xslices.Map(professionalsIDs, shared.StringToNanoid)

	a, err := w.repo.FindByWeekServiceAndProfessionals(date.Date(adate), nanoid.ID(serviceID), pids)
	if err != nil {
		return nil, err
	}

	return xslices.Map(a, toAppointmentOutput), nil
}
