package item

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/ajdinahmetovic/go-rest/db"
	"github.com/ajdinahmetovic/go-rest/httputil"
)

//GetItem func
func GetItem(w http.ResponseWriter, r *http.Request) {
	httputil.EnableCors(&w)
	v := r.URL.Query()
	id := v.Get("id")
	title := v.Get("title")
	description := v.Get("description")

	if id == "" && title == "" && description == "" {
		httputil.WriteResponse(w, "All items", *db.GetAllItems())
		return
	}

	data := *db.GetAllItems()
	var filter []db.Item
	/*
		if _id, err := strconv.Atoi(id); err != nil {
			for _, i := range data {
				if i.ID == _id {
					filter = append(filter, i)
				}
			}
		}
	*/

	_id, err := strconv.Atoi(id)
	if err != nil {
		_id = -1
	}

	for _, i := range data {
		if strings.HasPrefix(i.Title, title) && strings.HasPrefix(i.Description, description) && (i.ID == _id || _id == -1) {
			filter = append(filter, i)
		}
	}

	httputil.WriteResponse(w, "Items found", filter)

	/*
		_id, err := strconv.Atoi(id)
		if err != nil {

			return
		}
		item := db.FindItem(_id)

		if item != nil {
			httputil.WriteResponse(w, "Item found", &item)
		} else {
			httputil.WriteError(w, "Item NOT found", http.StatusNotFound)
		}
	*/

}
