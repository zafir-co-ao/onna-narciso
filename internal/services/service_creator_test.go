package services_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/duration"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

func TestServiceCreator(t *testing.T) {
	bus := event.NewEventBus()
	repo := services.NewInmemRepository()
	u := services.NewServiceCreator(repo, bus)

	t.Run("should_create_an_service", func(t *testing.T) {
		i := services.ServiceCreatorInput{
			Name:     "Manicure",
			Price:    "1000",
			Duration: 60,
		}

		_, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("must_save_the_service_in_the_repository", func(t *testing.T) {
		i := services.ServiceCreatorInput{
			Name:     "Massagem",
			Price:    "500",
			Duration: 120,
		}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		_, err = repo.FindByID(nanoid.ID(o.ID))
		if errors.Is(err, services.ErrServiceNotFound) {
			t.Errorf("Should return a service from repository, got %v", err)
		}
	})

	t.Run("must_register_the_duration_of_service", func(t *testing.T) {
		i := services.ServiceCreatorInput{
			Name:     "Pedicure",
			Price:    "1500",
			Duration: 60,
		}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if o.Duration != i.Duration {
			t.Errorf("The duration of service must be equal to %v, got %v", i.Duration, o.Duration)
		}
	})

	t.Run("must_register_the_name_of_service", func(t *testing.T) {
		i := services.ServiceCreatorInput{
			Name:     "Manicure",
			Price:    "2000",
			Duration: 150,
		}

		o, err := u.Create(i)
		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if o.Name != i.Name {
			t.Errorf("The name of service must be equal to %v, got %v", i.Name, o.Name)
		}
	})

	t.Run("must_register_the_price_of_service", func(t *testing.T) {
		i := services.ServiceCreatorInput{
			Name:        "Depilação",
			Description: "Depilação a laser",
			Price:       "1000",
			Duration:    120,
		}

		o, err := u.Create(i)
		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if o.Price != i.Price {
			t.Errorf("The price of service must be equal to %v, got %v", i.Price, o.Price)
		}
	})

	t.Run("must_register_the_description_of_service", func(t *testing.T) {
		i := services.ServiceCreatorInput{
			Name:        "Manicure",
			Description: "",
			Price:       "1000",
			Duration:    150,
		}

		o, err := u.Create(i)
		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if o.Description != i.Description {
			t.Errorf("The description of service must be equal to %v, got %v", i.Description, o.Description)
		}
	})

	t.Run("must_register_the_discount_of_service", func(t *testing.T) {
		i := services.ServiceCreatorInput{
			Name:        "Massagem",
			Description: "",
			Price:       "2500",
			Discount:    "10",
			Duration:    60,
		}

		o, err := u.Create(i)
		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if o.Discount != i.Discount {
			t.Errorf("The discount of service must be equal to %v, got %v", i.Discount, o.Discount)
		}
	})

	t.Run("the_duration_of_service_must_be_90_minutes_when_duration_not_provided", func(t *testing.T) {
		i := services.ServiceCreatorInput{
			Name:  "Manicure",
			Price: "1100",
		}

		o, err := u.Create(i)
		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(nanoid.ID(o.ID))
		if errors.Is(err, services.ErrServiceNotFound) {
			t.Errorf("Should return a service from repository, got %v", err)
		}

		if s.Duration != duration.DefaultDuration {
			t.Errorf("The duration of service must be Default: %v, got %v", duration.DefaultDuration, s.Duration)
		}
	})

	t.Run("should_publish_the_domain_event_when_service_is_created", func(t *testing.T) {
		isPublished := false
		var h event.HandlerFunc = func(e event.Event) {
			if e.Name() == services.EventServiceCreated {
				isPublished = true
			}
		}

		i := services.ServiceCreatorInput{
			Name:  "Manicure",
			Price: "1050",
		}

		bus.Subscribe(services.EventServiceCreated, h)

		_, err := u.Create(i)
		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if !isPublished {
			t.Errorf("The %v should be published", services.EventServiceCreated)
		}
	})

	t.Run("should_return_error_if_name_of_service_is_empty", func(t *testing.T) {
		i := services.ServiceCreatorInput{
			Name:     "",
			Duration: 90,
		}

		_, err := u.Create(i)
		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, name.ErrEmptyName) {
			t.Errorf("The error must be %v, got %v", name.ErrEmptyName, err)
		}
	})

	t.Run("should_return_error_if_duration_is_less_than_zero", func(t *testing.T) {
		i := services.ServiceCreatorInput{
			Name:     "Manicure",
			Duration: -90,
		}

		_, err := u.Create(i)
		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, duration.ErrInvalidDuration) {
			t.Errorf("The error must be %v, got %v", duration.ErrInvalidDuration, err)
		}
	})

	t.Run("should_return_error_if_price_is_empty", func(t *testing.T) {
		i := services.ServiceCreatorInput{
			Name:  "Manicure",
			Price: "",
		}

		_, err := u.Create(i)
		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, services.ErrInvalidPrice) {
			t.Errorf("The error must be %v, got %v", services.ErrInvalidPrice, err)
		}
	})

	t.Run("should_return_error_if_discount_is_invalid", func(t *testing.T) {
		i := services.ServiceCreatorInput{
			Name:     "Manicure",
			Price:    "1000",
			Discount: "1000",
		}

		_, err := u.Create(i)
		if errors.Is(nil, err) {
			t.Errorf("Expected error, got %v", err)
		}

		if !errors.Is(err, services.ErrDiscountNotAllowed) {
			t.Errorf("The error must be %v, got %v", services.ErrDiscountNotAllowed, err)
		}
	})
}
