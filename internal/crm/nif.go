package crm

type Nif string

func (n Nif) String() string {
	return string(n)
}
