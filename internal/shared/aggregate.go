package shared

import "github.com/zafir-co-ao/onna-narciso/internal/shared/id"

type Aggregate interface {
	GetID() id.ID
}
