package main

import (
	"fmt"
	"net/http"

	"github.com/ajdinahmetovic/go-rest/httputil"
	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

// Item struct
type Item struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var items = [3]Item{

	{
		ID:          "0",
		Title:       "This is title 1",
		Description: "Description of item 1",
	},

	{
		ID:          "1",
		Title:       "This is title 2",
		Description: "Description of item 2",
	},

	{
		ID:          "2",
		Title:       "This is title 3",
		Description: "Description of item 3",
	},
}

func getItems(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	v := r.URL.Query()
	id := v.Get("id")
	title := v.Get("title")
	description := v.Get("description")

	//Valide query params
	if id == "" {
		httputil.WriteError(w, "ID missing")
		return
	}

	if title == "" {
		httputil.WriteError(w, "Title missing")
		return
	}

	if description == "" {
		httputil.WriteError(w, "Description missing")
		return
	}

	item := Item{
		ID:          id,
		Title:       title,
		Description: description,
	}

	for _, i := range items {
		if i == item {
			httputil.WriteResponse(w, "Item found")
			return
		}
	}

	httputil.WriteError(w, "Item NOT found")

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {

	fmt.Println("Server starting")
	router.HandleFunc("/item", getItems).Methods("GET")
	err := http.ListenAndServe(":3000", router)

	if err != nil {
		fmt.Println("Server failed", err)
	}
	fmt.Println("Server started")

}
