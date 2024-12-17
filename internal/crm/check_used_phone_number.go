package crm

func checkUsedPhoneNumber(customers []Customer, p PhoneNumber) bool {
	for _, c := range customers {
		if c.PhoneNumber == p {
			return true
		}
	}
	return false
}
