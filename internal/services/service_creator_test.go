package services_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/services/adapters/inmem"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/duration"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

func TestServiceCreator(t *testing.T) {
	bus := event.NewEventBus()
	repo := inmem.NewServiceRepository()
	u := services.NewServiceCreator(repo, bus)

	t.Run("should_create_an_service", func(t *testing.T) {
		i := services.ServiceInput{
			Name:     "Manicure",
			Duration: 60,
		}

		_, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("must_save_the_service_in_the_repository", func(t *testing.T) {
		i := services.ServiceInput{
			Name:     "Massagem",
			Duration: 120,
		}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		_, err = repo.FindByID(nanoid.ID(o.ID))
		if errors.Is(services.ErrServiceNotFound, err) {
			t.Errorf("Should return a service from repository, got %v", err)
		}
	})

	t.Run("must_register_the_duration_of_service", func(t *testing.T) {
		i := services.ServiceInput{
			Name:     "Pedicure",
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
		i := services.ServiceInput{
			Name:     "Manicure",
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

	t.Run("must_register_the_description_of_service", func(t *testing.T) {
		i := services.ServiceInput{
			Name:        "Manicure",
			Description: "",
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

	t.Run("the_duration_of_service_must_be_90_minutes_when_duration_not_provided", func(t *testing.T) {
		i := services.ServiceInput{
			Name: "Manicure",
		}

		o, err := u.Create(i)
		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		s, err := repo.FindByID(nanoid.ID(o.ID))
		if errors.Is(services.ErrServiceNotFound, err) {
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

		i := services.ServiceInput{
			Name: "Manicure",
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
		i := services.ServiceInput{
			Name:     "",
			Duration: 90,
		}

		_, err := u.Create(i)
		if errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if !errors.Is(name.ErrInvalidName, err) {
			t.Errorf("The error must be %v, got %v", name.ErrInvalidName, err)
		}
	})

	t.Run("should_return_error_if_duration_is_less_than_zero", func(t *testing.T) {
		i := services.ServiceInput{
			Name:     "Manicure",
			Duration: -90,
		}

		_, err := u.Create(i)
		if errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if !errors.Is(duration.ErrInvalidDuration, err) {
			t.Errorf("The error must be %v, got %v", duration.ErrInvalidDuration, err)
		}
	})
}
