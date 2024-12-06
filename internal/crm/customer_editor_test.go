package crm_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/adapters/inmem"
)

func TestCustomerEdit(t *testing.T) {

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

	t.Run("should_find_the_customer", func(t *testing.T) {
	i := crm.CustomerEditorInput{
			ID:          "1",
			Name:        "Paola Miguel",
			NIF:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paola123.oliveira@domain.com",
			PhoneNumber: "+244922000022",
		}

		err := u.Edit(i)

		if err != nil {
			t.Errorf("expected no erro, got %v", err)
		}

		if errors.Is(err, crm.ErrCustomerNotFound) {
			t.Errorf("should find a customer in repository")
		}

	})

	t.Run("should_edit_name_of_customer", func(t *testing.T) {
	i := crm.CustomerEditorInput{
			ID:          "1",
			Name:        "Paola Miguel",
			NIF:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paola123.oliveira@domain.com",
			PhoneNumber: "+244922000022",
		}

		err := u.Edit(i)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		f, err := repo.FindByID(c.ID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if f.Name.String() != i.Name {
			t.Errorf("The name of customer %s should equal to %s", f.Name.String(), i.Name)

		}

	})

	t.Run("should_edit_nif_of_customer", func(t *testing.T) {
	i := crm.CustomerEditorInput{
			ID:          "1",
			Name:        "Paola Miguel",
			NIF:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paola123.oliveira@domain.com",
			PhoneNumber: "+244922000022",
		}

		err := u.Edit(i)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		f, err := repo.FindByID(c.ID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if f.Nif.String() != i.NIF {
			t.Errorf("The nif of customer %s should equal to %s", f.Nif.String(), i.NIF)

		}
	})

	t.Run("should_edit_birthdate_of_customer", func(t *testing.T) {

	i := crm.CustomerEditorInput{
			ID:          "1",
			Name:        "Paola Miguel",
			NIF:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paola123.oliveira@domain.com",
			PhoneNumber: "+244922000022",
		}

		err := u.Edit(i)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		f, err := repo.FindByID(c.ID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if f.BirthDate.String() != i.BirthDate {
			t.Errorf("The birthdateof customer %s should equal to %s", f.BirthDate.String(), i.BirthDate)

		}
	})

	t.Run("should_edit_email_of_customer", func(t *testing.T) {
	i := crm.CustomerEditorInput{
			ID:          "1",
			Name:        "Paola Miguel",
			NIF:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paola123.oliveira@domain.com",
			PhoneNumber: "+244922000022",
		}

		err := u.Edit(i)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		f, err := repo.FindByID(c.ID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if f.Email.String() != i.Email {
			t.Errorf("The email of customer %s should equal to %s", f.Email.String(), i.Email)

		}
	})

	t.Run("should_edit_phonenumber_of_customer", func(t *testing.T) {
		i := crm.CustomerEditorInput{
			ID:          "1",
			Name:        "Paola Miguel",
			NIF:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paola123.oliveira@domain.com",
			PhoneNumber: "+244922000022",
		}

		err := u.Edit(i)

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		f, err := repo.FindByID(c.ID)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if f.PhoneNumber.String() != i.PhoneNumber {
			t.Errorf("The phonenumber of customer %s should equal to %s", f.PhoneNumber.String(), i.PhoneNumber)

		}
	})
}
