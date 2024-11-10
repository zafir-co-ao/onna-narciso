package shared

import (
	"github.com/kindalus/godx/pkg/nanoid"
)

type BaseRepositoryConstraint interface {
	Aggregate
}

type BaseRepository[T BaseRepositoryConstraint] struct {
	Data map[nanoid.ID]T
}

func NewBaseRepository[T BaseRepositoryConstraint](s ...T) BaseRepository[T] {
	r := BaseRepository[T]{
		Data: make(map[nanoid.ID]T),
	}

	for _, a := range s {
		r.Data[a.GetID()] = a
	}

	return r
}
