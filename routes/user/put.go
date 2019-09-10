package user

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
		httputil.WriteError(w, "Check your JSON structure", http.StatusInternalServerError)
		return
	}
	var user db.User
	err = json.Unmarshal(req, &user)
	if err != nil {
		httputil.WriteError(w, "Error", http.StatusInternalServerError)
		return
	}
	err = db.UpdateUser(&user)
	if err != nil {
		httputil.WriteError(w, "Error while updating user", http.StatusInternalServerError)
		return
	}
	httputil.WriteResponse(w, "Item updated successfully", nil)
}
