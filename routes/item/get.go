package item

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ajdinahmetovic/go-rest/httputil"
	"github.com/ajdinahmetovic/item-service/proto/v1"
	"google.golang.org/grpc"
)

//GetItem func
func GetItem(w http.ResponseWriter, r *http.Request) {
	httputil.EnableCors(&w)

	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
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
	httputil.WriteResponse(w, items.Message, &items.Item)
}
