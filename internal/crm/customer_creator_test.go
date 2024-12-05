package crm_test

import (
	"errors"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/crm"
)

func TestCustomerCreator(t *testing.T) {
	t.Run("should_create_a_customer", func(t *testing.T) {
		u := crm.NewCustomerCreator()

		err := u.Create()

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}
