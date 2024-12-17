package hr

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/name"
)

type Professional struct {
	ID   nanoid.ID
	Name name.Name
}
