package hour

import (
	"errors"
	"time"
)

var ErrInvalidHour = errors.New("Invalid hour")

type Hour string

func New(v string) (Hour, error) {
	_, err := time.Parse("15:04", v)

	if err != nil {
		return Hour(""), ErrInvalidHour
	}
	return Hour(v), nil
}

func (h Hour) String() string {
	return string(h)
}

func AsTime(h Hour) time.Time {
	t, _ := time.Parse("15:04", h.String())
	return t
}

func Now() Hour {
	return Hour(time.Now().Format("15:04"))
}
