package crm

type PhoneNumber string

func (p PhoneNumber) String() string {
	return string(p)
}
