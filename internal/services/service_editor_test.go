package services_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/event"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
	"github.com/zafir-co-ao/onna-narciso/internal/services/adapters/inmem"
)

func TestServiceEdit(t *testing.T) {

	t.Run("should_recovery_service_with_id", func(t *testing.T) {
		bus := event.NewEventBus()
		repo := inmem.NewServiceRepository()
		u := services.NewServiceCreator(repo, bus)
		e := services.NewServiceEditor(repo)

		i := services.ServiceCreatorInput{
			Name:     "Manicure",
			Price:    "1000",
			Duration: 60,
		}

		o, err := u.Create(i)

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		err = e.Edit(string(o.ID))

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

	})

}
