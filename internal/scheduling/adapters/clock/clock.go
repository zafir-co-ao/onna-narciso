package clock

import (
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
)

func New() scheduling.ClockFunc {
	return func() date.Date {
		return date.Today()
	}
}
