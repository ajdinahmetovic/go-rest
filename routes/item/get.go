package item

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ajdinahmetovic/go-rest/db"
	"github.com/ajdinahmetovic/go-rest/httputil"
)

//GetItem func
func GetItem(w http.ResponseWriter, r *http.Request) {
	httputil.EnableCors(&w)

	v := r.URL.Query()
	id, err := strconv.Atoi(v.Get("id"))
	if err != nil {
		id = 0
	}

	title := v.Get("title")
	description := v.Get("description")

	queryItem := db.Item{ID: id, Title: title, Description: description}
	item, err := db.FindItem(&queryItem)

	if err != nil {
		fmt.Println(err)
		httputil.WriteError(w, "Failed to get item", http.StatusInternalServerError)
		return
	}

	httputil.WriteResponse(w, "Items found", &item)
}
