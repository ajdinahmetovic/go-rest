package user

import (
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
		id = 0
	}
	username := v.Get("username")
	fullname := v.Get("fullname")
	queryUser := db.User{ID: id, Username: username, FullName: fullname}

	users, err := db.FindUser(&queryUser)

	if err != nil {
		httputil.WriteError(w, "Error getting users", http.StatusInternalServerError)
		return
	}
	httputil.WriteResponse(w, "Users fetched successfully", users)
}
