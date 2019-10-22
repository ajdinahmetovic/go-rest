package user

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ajdinahmetovic/go-rest/config"
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
	conn, err := grpc.Dial(config.AppCfg.ItemServiceURL, grpc.WithInsecure())
	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	client := proto.NewUserServiceClient(conn)
	res, err := client.DeleteUser(context.Background(), &proto.DeleteUserReq{ID: int32(id)})
	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	httputil.WriteResponse(w, res.Message, nil)
}
