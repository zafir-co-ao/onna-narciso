package crm_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/crm"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

func TestCustomerCreator(t *testing.T) {
	bus := event.NewEventBus()
	repo := crm.NewInmemRepository()
	u := crm.NewCustomerCreator(repo, bus)

	t.Run("should_create_a_customer", func(t *testing.T) {
		i := crm.CustomerCreatorInput{
			Name:        "Paola Oliveira",
			Nif:         "002223109LA033",
			BirthDate:   "2000-01-02",
			Email:       "paola.oliveira@domain.com",
			PhoneNumber: "+244911000022",
		}

		_, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("should_save_customer_in_repository", func(t *testing.T) {
		i := crm.CustomerCreatorInput{
			Name:        "Joana Doe",
			Nif:         "002223109LA023",
			BirthDate:   "2005-05-10",
			Email:       "paola.oliveira1@domain.com",
			PhoneNumber: "+244932221100",
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
		i := crm.CustomerCreatorInput{
			Name:        "John Doe",
			Nif:         "002223109LA022",
			BirthDate:   "2006-01-01",
			Email:       "john.doe1@domain.com",
			PhoneNumber: "+244911112233",
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
		i := crm.CustomerCreatorInput{
			Name:        "Juliana Paes",
			Nif:         "002223109LA021",
			BirthDate:   "2006-08-12",
			Email:       "julianapaes@domain.com",
			PhoneNumber: "+244911909010",
		}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if o.Nif != i.Nif {
			t.Errorf("The nif of customer %s should equal to %s", i.Nif, o.Nif)
		}
	})

	t.Run("must_register_customer_without_the_birth_date", func(t *testing.T) {
		i := crm.CustomerCreatorInput{
			Name:        "Juliana Paes",
			Nif:         "002223109LA020",
			BirthDate:   "",
			Email:       "juliana.paes@domain.com",
			PhoneNumber: "+244922002324",
		}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if o.BirthDate != i.BirthDate {
			t.Errorf("The birth date of customer %s should equal to %s", i.BirthDate, o.BirthDate)
		}
	})

	t.Run("must_register_customer_without_the_email", func(t *testing.T) {
		i := crm.CustomerCreatorInput{
			Name:        "Juliana Paes",
			Nif:         "002223109LA034",
			BirthDate:   "1990-01-01",
			Email:       "juliana.paes1998@domain.com",
			PhoneNumber: "+244918888090",
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
		i := crm.CustomerCreatorInput{
			Name:        "Joana Doe",
			Nif:         "002223109LA031",
			BirthDate:   "1990-01-01",
			Email:       "joana.doe10@domain.com",
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

		i := crm.CustomerCreatorInput{
			Name:        "Paola Oliveira",
			Nif:         "002223109LA030",
			BirthDate:   "1990-01-01",
			Email:       "paola.oliveira1998@domain.com",
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
		i := crm.CustomerCreatorInput{
			Name:        "Juliana Paes",
			Nif:         "002223109LA033",
			BirthDate:   "1990-01-01",
			Email:       "paes@domain.com",
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

	t.Run("should_return_error_if_email_is_invalid", func(t *testing.T) {
		i := crm.CustomerCreatorInput{
			Name:        "Juliana Paes",
			Nif:         "002223109LA111",
			BirthDate:   "2001-10-15",
			Email:       "john.doe@",
			PhoneNumber: "+244912000011",
		}

		_, err := u.Create(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, crm.ErrInvalidEmailFormat) {
			t.Errorf("The error must be %v, got %v", crm.ErrInvalidEmailFormat, err)
		}
	})

	t.Run("should_return_error_if_name_of_customer_is_empty", func(t *testing.T) {
		i := crm.CustomerCreatorInput{
			Nif:         "002223109LA910",
			BirthDate:   "2001-10-15",
			Email:       "john.doe@domain.com",
			PhoneNumber: "+244912000011",
		}

		_, err := u.Create(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, name.ErrEmptyName) {
			t.Errorf("The error must be %v, got %v", name.ErrEmptyName, err)
		}
	})

	t.Run("should_return_error_if_nif_of_customer_is_empty", func(t *testing.T) {
		i := crm.CustomerCreatorInput{
			Name:      "Micheal Jordan",
			BirthDate: "2001-10-15",
			Email:     "michael.jordan@domain.com",
		}

		_, err := u.Create(i)

		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, crm.ErrEmptyNif) {
			t.Errorf("The error must be %v, got %v", crm.ErrEmptyNif, err)
		}
	})

	t.Run("should_return_error_if_birth_date_is_fomart_incorrect", func(t *testing.T) {})
}
