package auth

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared"
)

type inmemUserRepositoryImpl struct {
	shared.BaseRepository[User]
}

func NewInmemRepository(u ...User) Repository {
	return &inmemUserRepositoryImpl{BaseRepository: shared.NewBaseRepository[User](u...)}
}

func (r *inmemUserRepositoryImpl) FindByID(id nanoid.ID) (User, error) {
	if u, ok := r.Data[id]; ok {
		return u, nil
	}
	return User{}, ErrUserNotFound
}

func (r *inmemUserRepositoryImpl) FindByUsername(u Username) (User, error) {
	for _, user := range r.Data {
		if user.Username == u {
			return user, nil
		}
	}

	return User{}, ErrUserNotFound
}

func (r *inmemUserRepositoryImpl) Save(u User) error {
	r.Data[u.ID] = u
	return nil
}
