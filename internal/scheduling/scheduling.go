package scheduling

import "errors"

var (
	ErrBusyTime              = errors.New("Schedule time is busy")
	ErrCustomerNotFound      = errors.New("Client not found")
	ErrInvalidStatusToCancel = errors.New("Invalid status to cancel")
)
