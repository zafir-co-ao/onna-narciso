package crm

func checkUsedEmail(customers []Customer, v Email) bool {
	if isEmptyValue(v.String()) {
		return false
	}

	for _, c := range customers {
		if c.Email == v {
			return true
		}
	}
	return false
}

func checkUsedNif(customers []Customer, v Nif) bool {
	if isEmptyValue(v.String()) {
		return false
	}

	for _, c := range customers {
		if c.Nif == v {
			return true
		}
	}
	return false
}

func checkUsedPhoneNumber(customers []Customer, v PhoneNumber) bool {
	if isEmptyValue(v.String()) {
		return false
	}

	for _, c := range customers {
		if c.PhoneNumber == v {
			return true
		}
	}
	return false
}

func isEmptyValue(v string) bool {
	return len(v) == 0
}
