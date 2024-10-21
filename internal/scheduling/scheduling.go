package scheduling

import "errors"

var (
	ErrBusyTime              = errors.New("Schedule time is busy")
	ErrInvalidStatusToCancel = errors.New("Invalid status to cancel")
)
