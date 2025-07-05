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
	fmt.Fprintf(w, utils.FormatValue(cache[key].Value))
}

func SetItem(w http.ResponseWriter, r *http.Request) {
	cache := storage.GetCache()
	key := r.PathValue("key")

	var body types.RequestBody

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
