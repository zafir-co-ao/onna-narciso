package crm

func checkUsedNif(customers []Customer, v Nif) bool {
	for _, c := range customers {
		if c.Nif == v {
			return true
		}
	}
	return false
}
