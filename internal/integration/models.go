package integration

import "github.com/kindalus/godx/pkg/nanoid"

const EventSessionCheckedIn = "EventSessionCheckedIn"

type Appointment struct {
	ID     nanoid.ID
	Status string
}
