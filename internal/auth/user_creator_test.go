package auth_test

import (
	"errors"
	"testing"

	"github.com/zafir-co-ao/onna-narciso/internal/auth"
)


func TestUserCreator(t *testing.T) {
	t.Run("should_create_a_user", func(t *testing.T) {
		u := auth.NewUserCreator()

		err := u.Create()

		if !errors.Is(nil, err) {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}
