package item

import (
	"net/http"

	"github.com/ajdinahmetovic/go-rest/httputil"
)

// Item struct
type Item struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

//Items dummy array
var Items = [3]Item{

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

//Get func
func Get(w http.ResponseWriter, r *http.Request) {
	httputil.EnableCors(&w)
	v := r.URL.Query()
	id := v.Get("id")
	title := v.Get("title")
	description := v.Get("description")

	//Validate query params
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

	for _, i := range Items {
		if i == item {
			httputil.WriteResponse(w, "Item found")
			return
		}
	}

	httputil.WriteError(w, "Item NOT found")

}
