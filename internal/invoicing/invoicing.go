package invoicing

type Service struct {
	ID   string
	Name string
}

type InvoicingInput struct {
	CustomerID   string
	CustomerName string
	CustomerNIF  string
	Services     []Service
}

type Invoicing interface {
	Issue(i InvoicingInput) error
}

type InvoicingFunc func(i InvoicingInput) error

func (f InvoicingFunc) Issue(i InvoicingInput) error {
	return f(i)
}
