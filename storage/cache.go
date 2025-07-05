package storage

import (
	"fredis/types"
)

var cache = make(map[string]types.Item)

func GetCache() map[string]types.Item {
	return cache
}
