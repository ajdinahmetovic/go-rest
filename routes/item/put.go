package item

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ajdinahmetovic/go-rest/db"
	"github.com/ajdinahmetovic/go-rest/httputil"
)

//Put func
func Put(w http.ResponseWriter, r *http.Request) {
	httputil.EnableCors(&w)

	req, err := ioutil.ReadAll(r.Body)

	if err != nil {
		httputil.WriteError(w, "Failed to read body", http.StatusInternalServerError)
		return
	}

	var item db.Item
	err = json.Unmarshal(req, &item)
	if err != nil {
		httputil.WriteError(w, "Check your JSON structure", http.StatusNotAcceptable)
		return
	}

	err = db.UpdateItem(&item)
	if err != nil {
		httputil.WriteError(w, "Falied to update item", http.StatusInternalServerError)
		return
	}

	httputil.WriteResponse(w, "Item updated successfully", nil)

}
