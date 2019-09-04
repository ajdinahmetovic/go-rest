package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Item struct
type Item struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func getItems(w http.ResponseWriter, R *http.Request) {
	enableCors(&w)
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	item := 5
	fmt.Println(item)
	err := json.NewEncoder(w).Encode(item)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("djes")
	router.HandleFunc("/user", getItems).Methods("GET")
	http.ListenAndServe(":3000", router)
}
