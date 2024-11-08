package shared

import "github.com/zafir-co-ao/onna-narciso/internal/shared/id"

type BaseRepositoryConstraint interface {
	Aggregate
}

type BaseRepository[T BaseRepositoryConstraint] struct {
	Data map[id.ID]T
}

func NewBaseRepository[T BaseRepositoryConstraint](s ...T) BaseRepository[T] {
	r := BaseRepository[T]{
		Data: make(map[id.ID]T),
	}

	for _, a := range s {
		r.Data[a.GetID()] = a
	}

	return r
}
