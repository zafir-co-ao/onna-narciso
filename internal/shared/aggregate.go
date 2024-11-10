package shared

import (
	"github.com/kindalus/godx/pkg/nanoid"
)

type Aggregate interface {
	GetID() nanoid.ID
}
