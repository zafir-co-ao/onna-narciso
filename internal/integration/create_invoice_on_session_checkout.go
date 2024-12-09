package integration

import "github.com/kindalus/godx/pkg/event"

func NewCreateInvoiceOnSessionCheckoutListener() event.HandlerFunc {
	return func(e event.Event) {
		// Verificar se o evento é o certo.
		// assert
		
		assrt.NotNil(e)ert.NotNil(e)		
			}
}
