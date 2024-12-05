package crm_test

import (
	"errors"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/adapters/inmem"
)

func TestCustomerGetter(t *testing.T) {
	c := crm.Customer{ID: "1"}

	repo := inmem.NewCustomerRepository(c)
	u := crm.NewCustomerGetter(repo)

	t.Run("should_retrieve_a_customer_in_repository", func(t *testing.T) {
		o, err := u.Get(c.ID.String())

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error in get customer, got %v", err)
		}

		if o.ID != c.ID.String() {
			t.Errorf("should get customer with id %v, got %v", c.ID.String(), o.ID)
		}
	})

	t.Run("should_return_error_when_customer_not_exists_in_repository", func(t *testing.T) {
		_, err := u.Get("2")

		if errors.Is(nil, err) {
			t.Error("Expected error in get customer in repository")
		}

		if !errors.Is(err, crm.ErrCustomerNotFound) {
			t.Errorf("The error expeted must be %v, got %v", crm.ErrCustomerNotFound, err)
		}
	})
}
