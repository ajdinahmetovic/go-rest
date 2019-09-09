package item

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ajdinahmetovic/go-rest/db"
	"github.com/ajdinahmetovic/go-rest/httputil"
)

//Post func
func Post(w http.ResponseWriter, r *http.Request) {

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

	//item.ID = db.GetLenght()

	err = db.AddItem(&item)
	if err != nil {
		httputil.WriteError(w, "Error while saving item", http.StatusInternalServerError)
	}

	httputil.WriteResponse(w, "Saved succesfully", nil)

}
