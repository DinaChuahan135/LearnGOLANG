// main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items map[string]Item

func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	items[newItem.ID] = newItem
	w.WriteHeader(http.StatusCreated)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	itemID := r.URL.Query().Get("id")
	item, found := items[itemID]
	if !found {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(item)
}
func updateItem(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	updateditem, found := items[ID]
	if !found {
		http.Error(w, "Item not found", http.StatusNotFound)
		return

	} else {
		err := json.NewDecoder(r.Body).Decode(&updateditem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		items[updateditem.ID] = updateditem
		w.WriteHeader(http.StatusOK)
	}

}
func deleteItem(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	_, found := items[ID]
	if !found {
		http.Error(w, "Item not found", http.StatusNotFound)
		return

	} else {
		delete(items, ID)
		w.WriteHeader(http.StatusNoContent)
	}
}

func getAll(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(items)
}
func main() {
	items = make(map[string]Item)

	http.HandleFunc("/additem", createItem)
	http.HandleFunc("/item", getItem)
	http.HandleFunc("/allitem", getAll)
	http.HandleFunc("/deleteitem", deleteItem)
	http.HandleFunc("/updateitem", updateItem)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
