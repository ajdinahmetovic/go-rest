package user

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ajdinahmetovic/go-rest/config"
	"github.com/ajdinahmetovic/go-rest/db"
	"github.com/ajdinahmetovic/go-rest/httputil"
	"github.com/ajdinahmetovic/item-service/proto/v1"
	"google.golang.org/grpc"
)

//Put func
func Put(w http.ResponseWriter, r *http.Request) {
	httputil.EnableCors(&w, r)
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	var user db.User
	err = json.Unmarshal(req, &user)
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

	res, err := client.UpdateUser(context.Background(), &proto.UpdateUserReq{
		User: &proto.User{
			ID:       int32(user.ID),
			Username: user.Username,
			FullName: user.FullName,
		}})

	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	httputil.WriteResponse(w, res.Message, nil)
}
