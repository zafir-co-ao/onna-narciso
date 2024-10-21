package scheduling

import "time"

func VerifyAvailability(a Appointment, appointments []Appointment) bool {
	var isAvailable = true

	if len(appointments) == 0 {
		return isAvailable
	}

	for _, b := range appointments {
		if isNotAvailable(a, b) {
			isAvailable = false
			break
		}
	}

	return isAvailable

}

func isNotAvailable(a, b Appointment) bool {
	startTimeA, _ := time.Parse("15:04", a.Start)
	endTimeA, _ := time.Parse("15:04", a.End)
	startTimeB, _ := time.Parse("15:04", b.Start)
	endTimeB, _ := time.Parse("15:04", b.End)

	if startTimeA.Equal(startTimeB) {
		return true
	}

	if startTimeA.Before(startTimeB) && endTimeA.After(startTimeB) {
		return true
	}

	if startTimeA.After(startTimeB) && startTimeA.Before(endTimeB) {
		return true
	}

	return false
}
