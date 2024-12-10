package crm

type PhoneNumber string

func NewPhoneNumber(v string) (PhoneNumber, error) {
	return PhoneNumber(v), nil
}

func (p PhoneNumber) String() string {
	return string(p)
}
