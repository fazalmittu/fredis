package service

import (
	"encoding/json"
	"fmt"
	"fredis/storage"
	"fredis/types"
	"net/http"
)

// have to implement INCR, DECR, INCRBY, DECRBY

func CountController(w http.ResponseWriter, r *http.Request) {
	cache := storage.GetCache()
	key := r.PathValue("key")

	var body types.CounterRequestBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	item, exists := cache[key]
	if !exists {
		cache[key] = types.Item{
			Value: body.Amount,
		}
		fmt.Fprintf(w, "%d", body.Amount)
	} else {
		val := item.Value
		if v, ok := val.(int); ok {
			v += body.Amount
			item.Value = v
			cache[key] = item
			fmt.Fprintf(w, "%d", v)
		} else {
			http.Error(w, "Value is not an integer", http.StatusUnprocessableEntity)
		}
	}

	storage.Promote(key, item.Value)
}
