package scheduling

import (
	"time"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
)

func DateIsSpecification(d date.Date) shared.SpecificationFunc[Appointment] {
	return func(a Appointment) bool {
		return a.Date == d
	}
}

func StatusIsSpecification(s Status) shared.SpecificationFunc[Appointment] {
	return func(a Appointment) bool {
		return a.Status == s
	}
}

func NotCanceledIsSpecification() shared.SpecificationFunc[Appointment] {
	return func(a Appointment) bool {
		return a.Status != StatusCanceled
	}
}

func WeekIsSpecification(d date.Date) shared.SpecificationFunc[Appointment] {
	t1, err := time.Parse("2006-01-02", d.String())
	w1, y1 := t1.ISOWeek()

	if err != nil {
		return func(a Appointment) bool {
			return false
		}
	}

	return func(a Appointment) bool {
		t2, err := time.Parse("2006-01-02", a.Date.String())
		if err != nil {
			return false
		}

		w2, y2 := t2.ISOWeek()

		return w1 == w2 && y1 == y2
	}
}

func ServiceIsSpecification(sid nanoid.ID) shared.SpecificationFunc[Appointment] {
	return func(a Appointment) bool {
		return a.ServiceID.String() == sid.String()
	}
}

func ProfessionalsInSpecification(p []nanoid.ID) shared.SpecificationFunc[Appointment] {
	return func(a Appointment) bool {
		if len(p) == 0 {
			return true
		}

		for _, prof := range p {
			if a.ProfessionalID.String() == prof.String() {
				return true
			}
		}

		return false
	}
}

func ProfessionalIsSpecification(pid nanoid.ID) shared.SpecificationFunc[Appointment] {
	return func(a Appointment) bool {
		return a.ProfessionalID.String() == pid.String()
	}
}
