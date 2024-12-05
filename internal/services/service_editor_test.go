package services_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/services/adapters/inmem"
)

func TestServiceEdit(t *testing.T) {
	bus := event.NewEventBus()
	repo := inmem.NewServiceRepository()
	u := services.NewServiceCreator(repo, bus)
	e := services.NewServiceEditor(repo, bus)

	i := services.ServiceCreatorInput{
		Name:     "Manicure",
		Price:    "1000",
		Duration: 60,
	}

	t.Run("should_recovery_service_with_id", func(t *testing.T) {
		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		i := services.ServiceEditorInput{
			ID:       o.ID,
			Name:     "Manicure e Pedicure",
			Price:    "1500",
			Duration: 90,
		}

		err = e.Edit(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

	})

	t.Run("should_edit_name_of_service", func(t *testing.T) {
		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		i := services.ServiceEditorInput{
			ID:       o.ID,
			Name:     "Manicure e Pedicure",
			Price:    "1500",
			Duration: 90,
		}

		err = e.Edit(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("should_edit_price_of_service", func(t *testing.T) {
		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		i := services.ServiceEditorInput{
			ID:       o.ID,
			Name:     "Manicure e Pedicure",
			Price:    "1500",
			Duration: 90,
		}

		err = e.Edit(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("should_edit_description_of_service", func(t *testing.T) {
		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		i := services.ServiceEditorInput{
			ID:          o.ID,
			Name:        "Manicure e Pedicure",
			Price:       "1500",
			Description: "Com gelinho na no pé",
			Duration:    90,
		}

		err = e.Edit(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("should_edit_duration_of_service", func(t *testing.T) {
		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		i := services.ServiceEditorInput{
			ID:          o.ID,
			Name:        "Manicure e Pedicure",
			Price:       "1500",
			Description: "Com gelinho na no pé",
			Duration:    120,
		}

		err = e.Edit(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}
	})

	t.Run("should_publish_the_domain_event_when_service_is_edited", func(t *testing.T) {
		isPublished := false
		var h event.HandlerFunc = func(e event.Event) {
			if e.Name() == services.EventServiceEdited {
				isPublished = true
			}
		}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		i := services.ServiceEditorInput{
			ID:          o.ID,
			Name:        "Manicure e Pedicure",
			Price:       "1500",
			Description: "",
			Duration:    90,
		}

		bus.Subscribe(services.EventServiceEdited, h)

		err = e.Edit(i)
		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if !isPublished {
			t.Errorf("The %v should be published", services.EventServiceEdited)
		}
	})

}
