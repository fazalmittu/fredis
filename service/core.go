package service

import (
	"encoding/json"
	"fmt"
	"fredis/storage"
	"fredis/types"
	"fredis/utils"
	"net/http"
)

// have to implement SET, GET, DEL, EXISTS

func GetItem(w http.ResponseWriter, r *http.Request) {
	cache := storage.GetCache()
	key := r.PathValue("key")
	val, exists := cache[key]
	if !exists {
		http.Error(w, "nil", http.StatusNotFound)
		return
	}

	// promote *after* confirming existence
	storage.Promote(key, val.Value)

	fmt.Fprintf(w, "%s", utils.FormatValue(val.Value))

}

func SetItem(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	var body types.CoreRequestBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// cache[key] = types.Item{
	// 	Value: body.Value,
	// }
	storage.Promote(key, body.Value)

	fmt.Fprintf(w, "OK")

}

func DelItem(w http.ResponseWriter, r *http.Request) {
	cache := storage.GetCache()
	key := r.PathValue("key")

	_, exists := cache[key]

	if exists {
		// delete(cache, key)
		storage.Remove(key, true)
		fmt.Fprintf(w, "1")
	} else {
		fmt.Fprintf(w, "0")
	}
}

func Exists(w http.ResponseWriter, r *http.Request) {
	cache := storage.GetCache()
	key := r.PathValue("key")

	val, exists := cache[key]

	if exists {
		fmt.Fprintf(w, "1")
		storage.Promote(key, val.Value) // find a better way to do this
	} else {
		fmt.Fprintf(w, "0")
	}

}
