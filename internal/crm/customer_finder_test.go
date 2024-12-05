package crm_test

import (
	"errors"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/crm"
)

func TestCustomerFinder(t *testing.T) {
	t.Run("should find all customers", func(t *testing.T) {
		usecase := crm.NewCustomerFinder()

		o, err := usecase.Find()

		if !errors.Is(nil, err) {
			t.Error("Expected no error in find all customers")
		}

		if len(o) == 0 {
			t.Errorf("should find all customers, got %v", len(o))
		}
	})

}
