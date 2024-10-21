package scheduling

import (
	"errors"
	"fmt"
	"time"
)

var ErrInvalidHour = errors.New("Invalid hour")

type Hour string

func NewHour(v string) (Hour, error) {
	if !isValidHour(v) {
		return Hour(""), ErrInvalidHour
	}
	return Hour(v), nil
}

func (h Hour) Value() string {
	return string(h)
}

func isValidHour(v string) bool {
	x, err := time.Parse("15:04", v)

	fmt.Println(x)
	return err == nil
}
