package storage

import (
	"fredis/types"
)

var cache = make(map[string]types.Item)
var dll = types.DLL{}

func GetCache() map[string]types.Item {
	return cache
}

func GetDLL() types.DLL {
	return dll
}
