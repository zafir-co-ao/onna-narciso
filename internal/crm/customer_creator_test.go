package crm_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/crm/adapters/inmem"
)

func TestCustomerCreator(t *testing.T) {
	bus := event.NewEventBus()
	repo := inmem.NewCustomerRepository()
	u := crm.NewCustomerCreator(repo, bus)

	t.Run("should_create_a_customer", func(t *testing.T) {
		i := crm.CustomerCreatornput{
			Name: "Paola Oliveira",
			Nif:  "002223109LA033",
		}

		_, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("should_save_customer_in_repository", func(t *testing.T) {
		i := crm.CustomerCreatornput{
			Name: "Joana Doe",
			Nif:  "002223109LA023",
		}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		c, err := repo.FindByID(nanoid.ID(o.ID))
		if !errors.Is(nil, err) {
			t.Errorf("Expected no error when retrieve customer from repository, got %v", err)
		}

		if c.ID.String() != o.ID {
			t.Errorf("The id of customer %s should equal to %s", c.ID.String(), o.ID)
		}
	})

	t.Run("must_register_the_name_of_customer", func(t *testing.T) {
		i := crm.CustomerCreatornput{
			Name: "John Doe",
			Nif:  "002223109LA022",
		}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if o.Name != i.Name {
			t.Errorf("The name of customer %s should equal to %s", i.Name, o.Name)
		}
	})

	t.Run("must_register_the_nif_of_customer", func(t *testing.T) {
		i := crm.CustomerCreatornput{
			Name: "Juliana Paes",
			Nif:  "002223109LA021",
		}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if o.Nif != i.Nif {
			t.Errorf("The nif of customer %s should equal to %s", i.Nif, o.Nif)
		}
	})

	t.Run("must_register_the_birth_date_of_customer", func(t *testing.T) {
		i := crm.CustomerCreatornput{
			Name:      "Juliana Paes",
			Nif:       "002223109LA020",
			BirthDate: "1990-01-01",
		}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if o.BirthDate != i.BirthDate {
			t.Errorf("The birth date of customer %s should equal to %s", i.BirthDate, o.BirthDate)
		}
	})

	t.Run("must_register_the_email_of_customer", func(t *testing.T) {
		i := crm.CustomerCreatornput{
			Name:      "Juliana Paes",
			Nif:       "002223109LA034",
			BirthDate: "1990-01-01",
			Email:     "john.doe@domain.com",
		}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if o.Email != i.Email {
			t.Errorf("The email of customer %s should equal to %s", i.Email, o.Email)
		}
	})

	t.Run("must_register_the_phone_number_of_customer", func(t *testing.T) {
		i := crm.CustomerCreatornput{
			Name:        "Juliana Paes",
			Nif:         "002223109LA031",
			BirthDate:   "1990-01-01",
			Email:       "john.doe@domain.com",
			PhoneNumber: "+244912000011",
		}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if o.PhoneNumber != i.PhoneNumber {
			t.Errorf("The phoneNumber of customer %s should equal to %s", i.PhoneNumber, o.PhoneNumber)
		}
	})

	t.Run("should_publish_the_domain_event_when_customer_was_created", func(t *testing.T) {
		var isPublished bool = false

		i := crm.CustomerCreatornput{
			Name:        "Juliana Paes",
			Nif:         "002223109LA030",
			BirthDate:   "1990-01-01",
			Email:       "john.doe@domain.com",
			PhoneNumber: "+244912000011",
		}

		var h event.HandlerFunc = func(e event.Event) {
			if e.Name() == crm.EventCustomerCreated {
				isPublished = true
			}
		}

		bus.Subscribe(crm.EventCustomerCreated, h)

		_, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if !isPublished {
			t.Errorf("The %s must be published", crm.EventCustomerCreated)
		}
	})

	t.Run("should_return_error_if_nif_already_exists_in_repository", func(t *testing.T) {
		i := crm.CustomerCreatornput{
			Name:        "Juliana Paes",
			Nif:         "002223109LA033",
			BirthDate:   "1990-01-01",
			Email:       "john.doe@domain.com",
			PhoneNumber: "+244912000011",
		}

		_, err := u.Create(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, crm.ErrNifAlreadyUsed) {
			t.Errorf("The error mus be %v, got %v", crm.ErrNifAlreadyUsed, err)
		}
	})
}
