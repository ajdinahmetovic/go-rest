package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ajdinahmetovic/go-rest/db"
	"github.com/ajdinahmetovic/go-rest/httputil"
)

//Post func
func Post(w http.ResponseWriter, r *http.Request) {

	httputil.EnableCors(&w)

	req, err := ioutil.ReadAll(r.Body)

	if err != nil {
		httputil.WriteError(w, "Failed to read body", http.StatusInternalServerError)
		return
	}

	var user db.User
	err = json.Unmarshal(req, &user)
	if err != nil {
		httputil.WriteError(w, "Check your JSON", http.StatusInternalServerError)
		return
	}

	err = db.AddUser(&user)
	if err != nil {
		httputil.WriteError(w, "Failed to save user", http.StatusInternalServerError)
		return
	}
	httputil.WriteResponse(w, "User created successfully", nil)
}
