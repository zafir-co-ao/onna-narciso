package crm_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/adapters/inmem"
)

func TestCustomerEdit(t *testing.T) {

	t.Run("should_find_the_customer", func(t *testing.T) {
		c := crm.Customer{
			ID:          nanoid.ID("1"),
			Name:        "Paola Oliveira",
			Nif:         "002223109LA033",
			BirthDate:   "2000-01-02",
			Email:       "paola.oliveira@domain.com",
			PhoneNumber: "+244911000022",
		}

		repo := inmem.NewCustomerRepository(c)
		u := crm.NewCustomerEditor(repo)
		newName := "Paola Miguel"

		err := u.Edit(c.ID.String(), newName)

		if err != nil {
			t.Errorf("expected no erro, got %v", err)
		}

		if errors.Is(err, crm.ErrCustomerNotFound) {
			t.Errorf("should find a customer in repository")
		}

	})

	t.Run("should_edit_name_of_customer", func(t *testing.T) {
		c := crm.Customer{
			ID:          nanoid.ID("1"),
			Name:        "Paola Oliveira",
			Nif:         "002223109LA033",
			BirthDate:   "2000-01-02",
			Email:       "paola.oliveira@domain.com",
			PhoneNumber: "+244911000022",
		}

		repo := inmem.NewCustomerRepository(c)
		u := crm.NewCustomerEditor(repo)

		newName := "Paola Miguel"

		err := u.Edit(c.ID.String(), newName)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		f, err := repo.FindByID(c.ID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if f.Name != c.Name {
			t.Errorf("The name of customer %s should equal to %s", c.Name, f.Name)

		}

	})
}
