package user

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ajdinahmetovic/go-rest/config"
	"github.com/ajdinahmetovic/go-rest/httputil"
	"github.com/ajdinahmetovic/item-service/db"
	"github.com/ajdinahmetovic/item-service/proto/v1"
	"google.golang.org/grpc"
)

//Login func
func Login(w http.ResponseWriter, r *http.Request) {
	httputil.EnableCors(&w, r)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	var userCredidential db.UserCredentials
	err = json.Unmarshal(req, &userCredidential)
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
	res, err := client.Login(context.Background(), &proto.LoginUserReq{
		UserCredidentials: &proto.UserCredidentials{
			Username: userCredidential.Username,
			Password: userCredidential.Password,
		}})
	if err != nil {
		httputil.WriteError(w, err, http.StatusUnauthorized)
		return
	}
	httputil.WriteResponse(w, res.Message, map[string]string{
		"access_token":  res.Token,
		"refresh_token": res.RefreshToken,
		"user_id":       strconv.Itoa(int(res.UserID)),
	})
}
