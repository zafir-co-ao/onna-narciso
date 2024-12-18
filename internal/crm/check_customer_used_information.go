package crm

func checkUsedEmail(customers []Customer, v Email) bool {
	for _, c := range customers {
		if c.Email == v {
			return true
		}
	}
	return false
}

func checkUsedNif(customers []Customer, v Nif) bool {
	for _, c := range customers {
		if c.Nif == v {
			return true
		}
	}
	return false
}

func checkUsedPhoneNumber(customers []Customer, p PhoneNumber) bool {
	for _, c := range customers {
		if c.PhoneNumber == p {
			return true
		}
	}
	return false
}
