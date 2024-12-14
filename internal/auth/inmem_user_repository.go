package auth

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
)

type inmemUserRepository struct {
	shared.BaseRepository[User]
}

func NewInmemRepository(u ...User) Repository {
	return &inmemUserRepository{BaseRepository: shared.NewBaseRepository[User](u...)}
}

func (r *inmemUserRepository) FindByID(id nanoid.ID) (User, error) {
	if u, ok := r.Data[id]; ok {
		return u, nil
	}
	return User{}, ErrUserNotFound
}

func (r *inmemUserRepository) FindByUsername(u Username) (User, error) {
	for _, user := range r.Data {
		if user.Username == u {
			return user, nil
		}
	}

	return User{}, ErrUserNotFound
}

func (r *inmemUserRepository) Save(u User) error {
	r.Data[u.ID] = u
	return nil
}
