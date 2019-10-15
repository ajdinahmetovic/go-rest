package item

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ajdinahmetovic/go-rest/httputil"
	"github.com/ajdinahmetovic/item-service/proto/v1"
	"google.golang.org/grpc"
)

//Delete func
func Delete(w http.ResponseWriter, r *http.Request) {
	httputil.EnableCors(&w, r)
	v := r.URL.Query()
	id, err := strconv.Atoi(v.Get("id"))
	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	conn, err := grpc.Dial("service:4040", grpc.WithInsecure())
	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	client := proto.NewUserServiceClient(conn)

	res, err := client.DeleteItem(context.Background(), &proto.DeleteItemReq{ID: int32(id)})
	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	httputil.WriteResponse(w, res.Message, nil)

}
