package scheduling

import (
	"errors"
	"time"
)

var ErrInvalidDate = errors.New("Invalid date")

type Date string

func NewDate(v string) (Date, error) {
	_, err := time.Parse("2006-01-02", v)
	if err != nil {
		return Date(""), ErrInvalidDate
	}

	return Date(v), nil
}

func (d Date) Value() string {
	return string(d)
}
