package stubs

import (
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
)

type clock struct{}

func NewClock() scheduling.Clock {
	return &clock{}
}

func (c *clock) Today() date.Date {
	return date.Date("2016-01-01")
}
