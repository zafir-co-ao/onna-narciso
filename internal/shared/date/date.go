package date

import (
	"errors"
	"time"
)

var ErrInvalidDate = errors.New("Invalid date")

type Date string

func New(v string) (Date, error) {
	if !isValidDate(v) {
		return Date(""), ErrInvalidDate
	}

	return Date(v), nil
}

func (d Date) String() string {
	return string(d)
}

func isValidDate(v string) bool {
	_, err := time.Parse("2006-01-02", v)
	return err == nil
}
