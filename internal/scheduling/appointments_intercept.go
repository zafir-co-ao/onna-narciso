package scheduling

import (
	"time"

	"github.com/zafir-co-ao/onna-narciso/internal/shared/hour"
)

func AppointmentsIntersect(a, b Appointment) bool {

	if a.Hour.String() == b.Hour.String() {
		return true
	}

	atime := hour.AsTime(a.Hour)
	btime := hour.AsTime(b.Hour)

	if atime.After(btime) {
		atime, btime = btime, atime
		a, b = b, a
	}

	ctime := atime.Add(time.Duration(a.Duration) * time.Minute)

	return atime.Before(btime) && ctime.After(btime)
}

func AppointmentsInterceptAny(a Appointment, b []Appointment) bool {
	for _, appointment := range b {
		if AppointmentsIntersect(a, appointment) {
			return true
		}
	}
	return false
}
