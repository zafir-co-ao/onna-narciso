package crm_test

import (
	"errors"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/crm"
)

func TestCustomerFinder(t *testing.T) {
	t.Run("should_find_all_customers", func(t *testing.T) {
		customers := []crm.Customer{
			{
				ID:   "1",
				Name: "Jonatahan",
				Nif:  "08724GH3",
			},
			{
				ID:   "2",
				Name: "James",
				Nif:  "08724LH4",
			},
		}

		repo := crm.NewInmemRepository(customers...)
		u := crm.NewCustomerFinder(repo)

		o, err := u.Find()

		if !errors.Is(nil, err) {
			t.Error("Expected no error in find all customers")
		}

		if len(o) != 2 {
			t.Errorf("should find all customers, got %v", len(o))
		}
	})
}
