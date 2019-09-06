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

	data := *db.GetAllItems()
	var filter []db.Item
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
}
