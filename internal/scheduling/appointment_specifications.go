package scheduling

import (
	"time"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
)

func DateIsSpecificantion(d Date) shared.SpecificationFunc[Appointment] {
	return func(a Appointment) bool {
		return a.Date == d
	}
}

func StatusIsSpecificantion(s Status) shared.SpecificationFunc[Appointment] {
	return func(a Appointment) bool {
		return a.Status == s
	}
}

func WeekIsSpecificantion(d Date) shared.SpecificationFunc[Appointment] {
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

func ServiceIsSpecificantion(sid nanoid.ID) shared.SpecificationFunc[Appointment] {
	return func(a Appointment) bool {
		return a.ServiceID.String() == sid.String()
	}
}

func ProfessionalsInSpecificantion(p []nanoid.ID) shared.SpecificationFunc[Appointment] {
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

func ProfessionalIsSpecificantion(pid nanoid.ID) shared.SpecificationFunc[Appointment] {
	return func(a Appointment) bool {
		return a.ProfessionalID.String() == pid.String()
	}
}
