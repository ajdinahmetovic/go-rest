package item

import (
	"context"
	"net/http"

	"github.com/ajdinahmetovic/go-rest/es"
	"github.com/ajdinahmetovic/go-rest/httputil"
)

//GetItem func
func GetItem(w http.ResponseWriter, r *http.Request) {
	httputil.EnableCors(&w, r)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	//Extract data from route
	v := r.URL.Query()
	title := v.Get("title")

	//GetItem from elasticsearch
	result, err := es.GetItem(context.Background(), "title", title)
	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	httputil.WriteResponse(w, "Search successfull", &result.Hits.Hits)
}
