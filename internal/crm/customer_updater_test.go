package crm_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

func TestCustomerUpdate(t *testing.T) {
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
			PhoneNumber: "+244921000022",
		},
	}

	bus := event.NewEventBus()
	repo := crm.NewInmemRepository(customers...)
	u := crm.NewCustomerUpdater(repo, bus)

	t.Run("should_retrieve_the_customer", func(t *testing.T) {
		i := crm.CustomerUpdaterInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paola.oliveira@domain.com",
			PhoneNumber: "+244911000022",
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no erro, got %v", err)
		}

		o, err := repo.FindByID(customers[0].ID)
		if errors.Is(err, crm.ErrCustomerNotFound) {
			t.Errorf("Should find a customer in repository, got %v", err)
		}

		if o.ID.String() != i.ID {
			t.Errorf("expected %v, got %v", i.ID, o.ID.String())
		}
	})

	t.Run("should_update_name_of_customer", func(t *testing.T) {
		i := crm.CustomerUpdaterInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paola.oliveira@domain.com",
			PhoneNumber: "+244911000022",
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no erro, got %v", err)
		}

		c, err := repo.FindByID(nanoid.ID(i.ID))
		if errors.Is(err, crm.ErrCustomerNotFound) {
			t.Errorf("Should find a customer in repository, got %v", err)
		}

		if c.Name.String() != i.Name {
			t.Errorf("The name of customer %s should equal to %s", c.Name.String(), i.Name)
		}
	})

	t.Run("should_update_nif_of_customer", func(t *testing.T) {
		i := crm.CustomerUpdaterInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paola.oliveira@domain.com",
			PhoneNumber: "+244911000022",
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no erro, got %v", err)
		}

		c, err := repo.FindByID(nanoid.ID(i.ID))
		if errors.Is(err, crm.ErrCustomerNotFound) {
			t.Errorf("Should find a customer in repository, got %v", err)
		}

		if c.Nif.String() != i.Nif {
			t.Errorf("The nif of customer %s should equal to %s", c.Nif.String(), i.Nif)

		}
	})

	t.Run("should_update_birth_date_of_customer", func(t *testing.T) {
		i := crm.CustomerUpdaterInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paola.oliveira@domain.com",
			PhoneNumber: "+244911000022",
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no erro, got %v", err)
		}

		c, err := repo.FindByID(nanoid.ID(i.ID))
		if errors.Is(err, crm.ErrCustomerNotFound) {
			t.Errorf("Should find a customer in repository, got %v", err)
		}

		if c.BirthDate.String() != i.BirthDate {
			t.Errorf("The birthdateof customer %s should equal to %s", c.BirthDate.String(), i.BirthDate)
		}
	})

	t.Run("should_update_email_of_customer", func(t *testing.T) {
		i := crm.CustomerUpdaterInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paola.oliveira@domain.com",
			PhoneNumber: "+244911000022",
		}

		err := u.Update(i)
		if !errors.Is(nil, err) {
			t.Errorf("Expected no erro, got %v", err)
		}

		c, err := repo.FindByID(nanoid.ID(i.ID))
		if errors.Is(err, crm.ErrCustomerNotFound) {
			t.Errorf("Should find a customer in repository, got %v", err)
		}

		if c.Email.String() != i.Email {
			t.Errorf("The email of customer %s should equal to %s", c.Email.String(), i.Email)
		}
	})

	t.Run("should_update_phone_number_of_customer", func(t *testing.T) {
		i := crm.CustomerUpdaterInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paola.oliveira@domain.com",
			PhoneNumber: "+244922000022",
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no erro, got %v", err)
		}

		c, err := repo.FindByID(nanoid.ID(i.ID))
		if errors.Is(err, crm.ErrCustomerNotFound) {
			t.Errorf("Should find a customer in repository, got %v", err)
		}

		if c.PhoneNumber.String() != i.PhoneNumber {
			t.Errorf("The phonenumber of customer %s should equal to %s", c.PhoneNumber.String(), i.PhoneNumber)
		}
	})

	t.Run("should_return_error_when_the_customer_nif_already_exists", func(t *testing.T) {
		i := crm.CustomerUpdaterInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "001123109LA033",
			BirthDate:   "2001-01-02",
			Email:       "paola.oliveira@domain.com",
			PhoneNumber: "+244911000022",
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, crm.ErrNifAlreadyUsed) {
			t.Errorf("The error must be %v, got %v", crm.ErrNifAlreadyUsed, err)
		}
	})

	t.Run("should_publish_the_domain_event_when_customer_updated", func(t *testing.T) {
		var isPublished bool = false

		i := crm.CustomerUpdaterInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paola.oliveira@domain.com",
			PhoneNumber: "+244911000022",
		}

		var h event.HandlerFunc = func(e event.Event) {
			if e.Name() == crm.EventCustomerUpdated {
				isPublished = true
			}
		}

		bus.Subscribe(crm.EventCustomerUpdated, h)

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no erro, got %v", err)
		}

		if !isPublished {
			t.Errorf("The %s must be published", crm.EventCustomerUpdated)
		}
	})

	t.Run("should_return_error_if_email_is_invalid", func(t *testing.T) {
		i := crm.CustomerUpdaterInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "p123.oliveira@domain",
			PhoneNumber: "+244911000022",
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, crm.ErrInvalidEmailFormat) {
			t.Errorf("The error must be %v, got %v", crm.ErrInvalidEmailFormat, err)
		}
	})

	t.Run("should_return_error_if_name_is_empty", func(t *testing.T) {
		i := crm.CustomerUpdaterInput{
			ID:          "1",
			Name:        "",
			Nif:         "002223109LA031",
			BirthDate:   "2001-01-02",
			Email:       "paola.oliveira@domain.com",
			PhoneNumber: "+244911000022",
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, name.ErrEmptyName) {
			t.Errorf("The error must be %v, got %v", name.ErrEmptyName, err)
		}
	})

	t.Run("should_return_error_if_customer_not_exists_in_repository", func(t *testing.T) {
		i := crm.CustomerUpdaterInput{
			ID:   "100",
			Name: "John Doe",
			Nif:  "0023221ME048",
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, crm.ErrCustomerNotFound) {
			t.Errorf("The error must be %v, got %v", crm.ErrCustomerNotFound, err)
		}
	})

	t.Run("should_return_error_if_customer_email_is_already_used_by_other_customer", func(t *testing.T) {
		i := crm.CustomerUpdaterInput{
			ID:    "1",
			Name:  "Paola Oliveira",
			Nif:   "002223109LA033",
			Email: "monica12@domain.com",
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, crm.ErrEmailAlreadyUsed) {
			t.Errorf("The error must be %v, got %v", crm.ErrEmailAlreadyUsed, err)
		}
	})

	t.Run("should_return_error_if_customer_phone_number_is_used_by_other_customer", func(t *testing.T) {
		i := crm.CustomerUpdaterInput{
			ID:          "1",
			Name:        "Paola Oliveira",
			Nif:         "002223109LA033",
			Email:       "paola.oliveira@domain.com",
			PhoneNumber: "+244921000022",
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, crm.ErrPhoneNumberAlreadyUsed) {
			t.Errorf("The error must be %v, got %v", crm.ErrPhoneNumberAlreadyUsed, err)
		}
	})

	t.Run("should_return_error_if_birth_date_is_lower_than_minimum_allowed_age", func(t *testing.T) {
		i := crm.CustomerUpdaterInput{
			ID:          "1",
			Name:        "Paola Miguel",
			Nif:         "002223109LA031",
			BirthDate:   "2015-01-02",
			Email:       "paola.oliveira@domain.com",
			PhoneNumber: "+244911000022",
		}

		err := u.Update(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, crm.ErrAgeNotAllowed) {
			t.Errorf("The error must be %v, got %v", crm.ErrAgeNotAllowed, err)
		}
	})

	t.Run("should_update_with_empty_email", func(t *testing.T) {
		i := crm.CustomerUpdaterInput{
			ID:          "1",
			Name:        "Paola Oliveira",
			Nif:         "002223109LA033",
			Email:       "",
			PhoneNumber: "+244911000022",
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		c, err := repo.FindByID(nanoid.ID(i.ID))
		if errors.Is(err, crm.ErrCustomerNotFound) {
			t.Errorf("Should find a customer in repository, got %v", err)
		}

		if c.Email.String() != i.Email {
			t.Errorf("The customer email should be empty, got %v", c.Email.String())
		}
	})

	t.Run("should_update_with_empty_phone_number", func(t *testing.T) {
		i := crm.CustomerUpdaterInput{
			ID:          "1",
			Name:        "Paola Oliveira",
			Nif:         "002223109LA033",
			Email:       "paola.oliveira@domain.com",
			PhoneNumber: "",
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		c, err := repo.FindByID(nanoid.ID(i.ID))
		if errors.Is(err, crm.ErrCustomerNotFound) {
			t.Errorf("Should find a customer in repository, got %v", err)
		}

		if c.PhoneNumber.String() != i.PhoneNumber {
			t.Errorf("The customer phone number should be empty, got %v", c.PhoneNumber.String())
		}
	})

	t.Run("should_update_with_empty_birth_date", func(t *testing.T) {
		i := crm.CustomerUpdaterInput{
			ID:          "1",
			Name:        "Paola Oliveira",
			BirthDate:   "",
			Nif:         "002223109LA033",
			Email:       "paola.oliveira@domain.com",
			PhoneNumber: "+244911000022",
		}

		err := u.Update(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		c, err := repo.FindByID(nanoid.ID(i.ID))
		if errors.Is(err, crm.ErrCustomerNotFound) {
			t.Errorf("Should find a customer in repository, got %v", err)
		}

		if c.BirthDate.String() != i.BirthDate {
			t.Errorf("The customer birth date should be empty, got %v", c.BirthDate.String())
		}
	})
}
