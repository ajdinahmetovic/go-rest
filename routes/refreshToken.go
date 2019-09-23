package routes

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ajdinahmetovic/go-rest/httputil"
	"github.com/ajdinahmetovic/item-service/proto/v1"
	"google.golang.org/grpc"
)

type refreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

//RefreshToken func
func RefreshToken(w http.ResponseWriter, r *http.Request) {
	httputil.EnableCors(&w)
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	var refreshRequest refreshRequest
	err = json.Unmarshal(req, &refreshRequest)
	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	client := proto.NewUserServiceClient(conn)

	res, err := client.RefreshToken(context.Background(), &proto.RefreshTokenReq{
		RefreshToken: refreshRequest.RefreshToken,
	})
	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	httputil.WriteResponse(w, "Token refresed", map[string]string{
		"refresh_toke": res.RefreshToken,
		"access_token": res.Token,
	})

}
