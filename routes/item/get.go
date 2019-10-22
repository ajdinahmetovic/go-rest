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

	//GetItem from elasticsearch
	result, err := es.GetItem(context.Background())
	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	/*
		//Search with grpc and postgress
		conn, err := grpc.Dial(config.AppCfg.ItemServiceURL, grpc.WithInsecure())
		if err != nil {
			httputil.WriteError(w, err, http.StatusInternalServerError)
			return
		}
		client := proto.NewUserServiceClient(conn)

		v := r.URL.Query()
		id, err := strconv.Atoi(v.Get("id"))
		if err != nil {
			id = 0
		}

		title := v.Get("title")
		description := v.Get("description")
		items, err := client.GetItem(context.Background(), &proto.GetItemReq{
			ID:          int32(id),
			Title:       title,
			Description: description,
			UserID:      int32(id),
		})
		if err != nil {
			httputil.WriteError(w, err, http.StatusInternalServerError)
			return
		}
	*/

	httputil.WriteResponse(w, "Search successfull", &result)
}
