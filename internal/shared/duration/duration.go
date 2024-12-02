package duration

import "errors"

const DefaultDuration = 90

var ErrInvalidDuration = errors.New("invalid duration")

type Duration int

func New(v int) (Duration, error) {
	if isLessThanZero(v) {
		return 0, ErrInvalidDuration
	}

	if isZero(v) {
		return Duration(DefaultDuration), nil
	}

	return Duration(v), nil
}

func (d Duration) Value() int {
	return int(d)
}

func isLessThanZero(v int) bool {
	return v < 0
}

func isZero(v int) bool {
	return v == 0
}
