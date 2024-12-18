package shared

import "github.com/kindalus/godx/pkg/nanoid"

func StringToNanoid(id string) nanoid.ID {
	return nanoid.ID(id)
}

func NanoidToString(id nanoid.ID) string {
	return id.String()
}
