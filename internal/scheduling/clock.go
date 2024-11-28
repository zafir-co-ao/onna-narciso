package scheduling

import "github.com/zafir-co-ao/onna-narciso/internal/shared/date"

type Clock interface {
	Today() date.Date
}

type ClockFunc func() date.Date

func (f ClockFunc) Today() date.Date {
	return f()
}
