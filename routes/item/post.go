package item

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ajdinahmetovic/go-rest/config"
	"github.com/ajdinahmetovic/go-rest/logger"

	"github.com/ajdinahmetovic/go-rest/db"
	"github.com/ajdinahmetovic/go-rest/es"
	"github.com/ajdinahmetovic/go-rest/httputil"
	"github.com/ajdinahmetovic/item-service/proto/v1"
	"google.golang.org/grpc"
)

//Post func
func Post(w http.ResponseWriter, r *http.Request) {
	httputil.EnableCors(&w, r)

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	var itemReq db.Item

	err = json.Unmarshal(req, &itemReq)
	if err != nil {
		httputil.WriteError(w, err, http.StatusNotAcceptable)
		return
	}

	conn, err := grpc.Dial(config.AppCfg.ItemServiceURL, grpc.WithInsecure())
	if err != nil {
		httputil.WriteResponse(w, "Connection to item service failed", http.StatusInternalServerError)
		return
	}

	item := &proto.Item{
		Title:       itemReq.Title,
		Description: itemReq.Description,
		UserID:      int32(itemReq.UserID),
	}

	client := proto.NewUserServiceClient(conn)
	res, err := client.CreateItem(context.Background(), &proto.CreateItemReq{Item: item})
	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	//Set DB id and insert in es
	item.ID = res.ID
	err = es.InsertItem(context.Background(), item)
	if err != nil {
		logger.Error("Failed to save item to index", "item", item, "error", err)
		httputil.WriteError(w, err, http.StatusInternalServerError)

		return
	}

	httputil.WriteResponse(w, res.Message, nil)
}
