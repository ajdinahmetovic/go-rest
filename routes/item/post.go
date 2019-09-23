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

//Post func
func Post(w http.ResponseWriter, r *http.Request) {
	httputil.EnableCors(&w)

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

	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		httputil.WriteResponse(w, "Connection to item service failed", http.StatusInternalServerError)
		return
	}

	client := proto.NewUserServiceClient(conn)

	msg, err := client.CreateItem(context.Background(), &proto.CreateItemReq{Item: &proto.Item{
		Title:       itemReq.Title,
		Description: itemReq.Description,
		UserID:      int32(itemReq.UserID),
	}})
	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	httputil.WriteResponse(w, msg.Message, nil)
}
