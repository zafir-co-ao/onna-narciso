package crm

func checkUsedEmail(customers []Customer, v Email) bool {
	for _, c := range customers {
		if c.Email == v {
			return true
		}
	}
	return false
}
