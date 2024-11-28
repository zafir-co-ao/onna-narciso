package console

import (
	"fmt"

	"github.com/zafir-co-ao/onna-narciso/internal/sessions"
)

func NewInvoicing() sessions.InvoicingFunc {
	return func(s sessions.Session) error {
		fmt.Println(fmt.Sprintf("Invoicing... to %s with services: %v", s.CustomerName, s.Services))
		return nil
	}
}
