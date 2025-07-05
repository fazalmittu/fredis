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
	fmt.Fprintf(w, "%s", utils.FormatValue(cache[key].Value)) // use format str bc func expects a constant as second param
}

func SetItem(w http.ResponseWriter, r *http.Request) {
	cache := storage.GetCache()
	key := r.PathValue("key")

	var body types.CoreRequestBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	cache[key] = types.Item{
		Value: body.Value,
	}

	fmt.Fprintf(w, "OK")

}

func DelItem(w http.ResponseWriter, r *http.Request) {
	cache := storage.GetCache()
	key := r.PathValue("key")

	_, exists := cache[key]

	if exists {
		delete(cache, key)
		fmt.Fprintf(w, "1")
	} else {
		fmt.Fprintf(w, "0")
	}
}

func Exists(w http.ResponseWriter, r *http.Request) {
	cache := storage.GetCache()
	key := r.PathValue("key")

	_, exists := cache[key]

	if exists {
		fmt.Fprintf(w, "1")
	} else {
		fmt.Fprintf(w, "0")
	}

}
