package stubs

import (
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
)

func NewClock() scheduling.ClockFunc {
	return func() date.Date {
		return date.Date("2016-01-01")
	}
}
