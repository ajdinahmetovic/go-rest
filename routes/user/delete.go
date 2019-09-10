package user

import (
	"net/http"
	"strconv"

	"github.com/ajdinahmetovic/go-rest/db"
	"github.com/ajdinahmetovic/go-rest/httputil"
)

//Delete func
func Delete(w http.ResponseWriter, r *http.Request) {
	httputil.EnableCors(&w)

	v := r.URL.Query()
	id, err := strconv.Atoi(v.Get("id"))

	if err != nil {
		httputil.WriteError(w, "Invalid ID", http.StatusInternalServerError)
		return
	}

	err = db.DeleteUser(id)
	if err != nil {
		httputil.WriteError(w, "Error while deleting user", http.StatusInternalServerError)
		return
	}

	httputil.WriteResponse(w, "Successfully deleted", nil)
}
