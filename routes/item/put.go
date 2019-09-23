package item

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ajdinahmetovic/go-rest/db"
	"github.com/ajdinahmetovic/go-rest/httputil"
	"github.com/ajdinahmetovic/item-service/proto/v1"
	"google.golang.org/grpc"
)

//Put func
func Put(w http.ResponseWriter, r *http.Request) {
	httputil.EnableCors(&w)
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	client := proto.NewUserServiceClient(conn)

	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	var item db.Item
	err = json.Unmarshal(req, &item)
	if err != nil {
		httputil.WriteError(w, err, http.StatusNotAcceptable)
		return
	}

	res, err := client.UpdateItem(context.Background(), &proto.UpdateItemReq{
		Item: &proto.Item{
			ID:          int32(item.ID),
			Title:       item.Title,
			Description: item.Description,
			UserID:      int32(item.UserID),
		}})
	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	httputil.WriteResponse(w, res.Message, nil)
}
