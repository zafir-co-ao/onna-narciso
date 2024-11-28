package sessions

type Invoicing interface {
	Issue(s Session) error
}

type InvoicingFunc func(s Session) error

func (f InvoicingFunc) Issue(s Session) error {
	return f(s)
}
