package item

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ajdinahmetovic/go-rest/db"
	"github.com/ajdinahmetovic/go-rest/httputil"
)

//Get func
func Get(w http.ResponseWriter, r *http.Request) {
	httputil.EnableCors(&w)
	v := r.URL.Query()
	id, err := strconv.Atoi(v.Get("id"))

	if err != nil {
		httputil.WriteError(w, "Invalid ID ", http.StatusNotFound)
		fmt.Println(err)
		return
	}

	title := v.Get("title")
	description := v.Get("description")

	//Validate query params
	if title == "" {
		httputil.WriteError(w, "Title missing", http.StatusNotFound)
		return
	}

	if description == "" {
		httputil.WriteError(w, "Description missing", http.StatusNotFound)
		return
	}

	item := db.Item{
		ID:          id,
		Title:       title,
		Description: description,
	}

	for _, i := range db.DATA {
		if i == item {
			httputil.WriteResponse(w, "Item found", item)
			return
		}
	}

	httputil.WriteError(w, "Item NOT found", http.StatusNotFound)

}
