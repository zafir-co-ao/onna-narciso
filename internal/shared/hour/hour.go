package hour

import (
	"errors"
	"time"
)

var ErrInvalidFormat = errors.New("invalid hour")

type Hour string

func New(v string) (Hour, error) {
	if !isValidFormat(v) {
		return Hour(""), ErrInvalidFormat
	}

	return Hour(v), nil
}

func (h Hour) String() string {
	return string(h)
}

func isValidFormat(v string) bool {
	_, err := time.Parse("15:04", v)
	return err == nil
}

func AsTime(h Hour) time.Time {
	t, _ := time.Parse("15:04", h.String())
	return t
}

func Now() Hour {
	return Hour(time.Now().Format("15:04"))
}
