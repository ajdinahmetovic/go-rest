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
	id, err := strconv.Atoi(v.Get("id"))
	if err != nil {
		id = -1
	}

	title := v.Get("title")
	description := v.Get("description")

	data := *db.GetAllItems()
	var filter []db.Item

	for _, i := range data {
		if strings.HasPrefix(i.Title, title) && strings.HasPrefix(i.Description, description) && (i.ID == id || id == -1) {
			filter = append(filter, i)
		}
	}

	httputil.WriteResponse(w, "Items found", filter)
}
