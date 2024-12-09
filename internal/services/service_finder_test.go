package services_test

import (
	"errors"
	"testing"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/services"
)

func TestServiceFinder(t *testing.T) {
	s1 := services.Service{ID: nanoid.New()}
	s2 := services.Service{ID: nanoid.New()}
	repo := services.NewInmemRepository(s1, s2)

	t.Run("should find all services in repository", func(t *testing.T) {

		u := services.NewServiceFinder(repo)

		o, err := u.Find()

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(o) < 2 {
			t.Errorf("Should return list of services, got %v", o)
		}
	})
}
