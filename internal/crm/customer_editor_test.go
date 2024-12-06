package crm_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/adapters/inmem"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/email"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/nif"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/phone"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

func TestCustomerEdit(t *testing.T) {
	customers := []crm.Customer{
		{
			ID:          nanoid.ID("1"),
			Name:        "Paola Oliveira",
			Nif:         "002223109LA033",
			BirthDate:   "2000-01-02",
			Email:       "paola.oliveira@domain.com",
			PhoneNumber: "+244911000022",
		},
		{
			ID:          nanoid.ID("2"),
			Name:        "Monica Carlos",
			Nif:         "001123109LA033",
			BirthDate:   "2000-01-02",
			Email:       "monica12@domain.com",
			PhoneNumber: "+244911000022",
		},
	}

	bus := event.NewEventBus()
	repo := inmem.NewCustomerRepository(customers...)
	u := crm.NewCustomerEditor(repo, bus)

	t.Run("should_find_the_customer", func(t *testing.T) {
		i := crm.CustomerEditorInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paola123.oliveira@domain.com",
			PhoneNumber: "+244922000022",
		}

		err := u.Edit(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no erro, got %v", err)
		}

		if errors.Is(err, crm.ErrCustomerNotFound) {
			t.Errorf("Should find a customer in repository, got %v", err)
		}

	})

	t.Run("should_update_name_of_customer", func(t *testing.T) {
		i := crm.CustomerEditorInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paola123.oliveira@domain.com",
			PhoneNumber: "+244922000022",
		}

		err := u.Edit(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no erro, got %v", err)
		}

		c, err := repo.FindByID(i.ID)
		if errors.Is(err, crm.ErrCustomerNotFound) {
			t.Errorf("Should find a customer in repository, got %v", err)
		}

		if c.Name.String() != i.Name {
			t.Errorf("The name of customer %s should equal to %s", c.Name.String(), i.Name)
		}

	})

	t.Run("should_update_nif_of_customer", func(t *testing.T) {
		i := crm.CustomerEditorInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paola123.oliveira@domain.com",
			PhoneNumber: "+244922000022",
		}

		err := u.Edit(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no erro, got %v", err)
		}

		c, err := repo.FindByID(i.ID)
		if errors.Is(err, crm.ErrCustomerNotFound) {
			t.Errorf("Should find a customer in repository, got %v", err)
		}

		if c.Nif.String() != i.Nif {
			t.Errorf("The nif of customer %s should equal to %s", c.Nif.String(), i.Nif)

		}
	})

	t.Run("should_update_birth_date_of_customer", func(t *testing.T) {
		i := crm.CustomerEditorInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paola123.oliveira@domain.com",
			PhoneNumber: "+244922000022",
		}

		err := u.Edit(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no erro, got %v", err)
		}

		c, err := repo.FindByID(i.ID)
		if errors.Is(err, crm.ErrCustomerNotFound) {
			t.Errorf("Should find a customer in repository, got %v", err)
		}

		if c.BirthDate.String() != i.BirthDate {
			t.Errorf("The birthdateof customer %s should equal to %s", c.BirthDate.String(), i.BirthDate)

		}
	})

	t.Run("should_update_email_of_customer", func(t *testing.T) {
		i := crm.CustomerEditorInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paola123.oliveira@domain.com",
			PhoneNumber: "+244922000022",
		}

		err := u.Edit(i)
		if !errors.Is(nil, err) {
			t.Errorf("Expected no erro, got %v", err)
		}

		c, err := repo.FindByID(i.ID)
		if errors.Is(err, crm.ErrCustomerNotFound) {
			t.Errorf("Should find a customer in repository, got %v", err)
		}

		if c.Email.String() != i.Email {
			t.Errorf("The email of customer %s should equal to %s", c.Email.String(), i.Email)

		}
	})

	t.Run("should_update_phone_number_of_customer", func(t *testing.T) {
		i := crm.CustomerEditorInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paola123.oliveira@domain.com",
			PhoneNumber: "+244922000022",
		}

		err := u.Edit(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no erro, got %v", err)
		}

		c, err := repo.FindByID(i.ID)
		if errors.Is(err, crm.ErrCustomerNotFound) {
			t.Errorf("Should find a customer in repository, got %v", err)
		}

		if c.PhoneNumber.String() != i.PhoneNumber {
			t.Errorf("The phonenumber of customer %s should equal to %s", c.PhoneNumber.String(), i.PhoneNumber)

		}
	})

	t.Run("should_return_error_when_the_customer_nif_already_exists", func(t *testing.T) {
		i := crm.CustomerEditorInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "001123109LA033",
			BirthDate:   "2001-01-02",
			Email:       "paola123.oliveira@domain.com",
			PhoneNumber: "+244922000022",
		}

		err := u.Edit(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, crm.ErrNifAlreadyUsed) {
			t.Errorf("The error must be %v, got %v", crm.ErrNifAlreadyUsed, err)
		}
	})

	t.Run("should_publish_the_domain_event_when_customer_was_updated", func(t *testing.T) {
		var isPublished bool = false

		i := crm.CustomerEditorInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paola123.oliveira@domain.com",
			PhoneNumber: "+244922000022",
		}

		var h event.HandlerFunc = func(e event.Event) {
			if e.Name() == crm.EventCustomerUpdated {
				isPublished = true
			}
		}

		bus.Subscribe(crm.EventCustomerUpdated, h)

		err := u.Edit(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no erro, got %v", err)
		}

		_, err = repo.FindByID(i.ID)
		if errors.Is(err, crm.ErrCustomerNotFound) {
			t.Errorf("Should find a customer in repository, got %v", err)
		}

		if !isPublished {
			t.Errorf("The %s must be published", crm.EventCustomerUpdated)
		}
	})

	t.Run("should_return_error_if_birth_date_is_invalid", func(t *testing.T) {
		i := crm.CustomerEditorInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "002223109LA031",
			BirthDate:   "2001/01/02",
			Email:       "paola123.oliveira@domain.com",
			PhoneNumber: "+244922000022",
		}

		err := u.Edit(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, date.ErrInvalidFormat) {
			t.Errorf("The error must be %v, got %v", crm.ErrNifAlreadyUsed, err)
		}
	})

	t.Run("should_return_error_if_email_is_invalide", func(t *testing.T) {
		i := crm.CustomerEditorInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "p123.oliveira@domain",
			PhoneNumber: "+244922000022",
		}

		err := u.Edit(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, email.ErrInvalidFormat) {
			t.Errorf("The error must be %V, got %V", email.ErrInvalidFormat, err)
		}
	})

	t.Run("should_return_error_if_name_is_empty", func(t *testing.T) {
		i := crm.CustomerEditorInput{
			ID:          "1",
			Name:        "",
			Nif:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paulaoliveira.oliveira@domain.com",
			PhoneNumber: "+244922000022",
		}

		err := u.Edit(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, name.ErrEmptyName) {
			t.Errorf("The error must be %V, got %V", name.ErrEmptyName, err)
		}
	})

	t.Run("should_return_error_if_phone_number_of_customer_is_empy", func(t *testing.T) {
		i := crm.CustomerEditorInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paulaoliveira.oliveira@domain.com",
			PhoneNumber: "",
		}

		err := u.Edit(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, phone.ErrEmptyPhoneNumber) {
			t.Errorf("The error must be %V, got %V", phone.ErrEmptyPhoneNumber, err)
		}
	})

	t.Run("should_return_error_if_nif_of_customer_is_empty", func(t *testing.T) {
		i := crm.CustomerEditorInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "",
			BirthDate:   "2001-01-02",
			Email:       "paulaoliveira.oliveira@domain.com",
			PhoneNumber: "244922000022",
		}

		err := u.Edit(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, nif.ErrEmptyNif) {
			t.Errorf("The error must be %V, got %V", nif.ErrEmptyNif, err)
		}
	})
}
