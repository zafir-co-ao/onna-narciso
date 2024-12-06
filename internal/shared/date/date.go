package date

import (
	"errors"
	"time"
)

var (
	ErrInvalidFormat = errors.New("Invalid format of date")
	ErrDateInPast    = errors.New("Date in past not allowed")
)

type Date string

func New(v string) (Date, error) {
	if !IsValidFormat(v) {
		return Date(""), ErrInvalidFormat
	}

	d := Date(v)

	return d, nil
}

func Today() Date {
	return Date(time.Now().Format("2006-01-02"))
}

func (d Date) String() string {
	return string(d)
}

func (d Date) Before() bool {
	a, _ := time.Parse("2006-01-02", d.String())
	b, _ := time.Parse("2006-01-02", Today().String())
	return a.Before(b)
}

func (d Date) AddDate(years, months, days int) Date {
	a, _ := time.Parse("2006-01-02", d.String())
	return Date(a.AddDate(years, months, days).Format("2006-01-02"))
}

func IsValidFormat(v string) bool {
	_, err := time.Parse("2006-01-02", v)
	return err == nil
}
