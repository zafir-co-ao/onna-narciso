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

func (d Date) Before(v Date) bool {
	a, _ := time.Parse("2006-01-02", d.String())
	b, _ := time.Parse("2006-01-02", v.String())
	return a.Before(b)
}

func Today() Date {
	return Date(time.Now().Format("2006-01-02"))
}

func isValidDate(v string) bool {
	_, err := time.Parse("2006-01-02", v)
	return err == nil
}
