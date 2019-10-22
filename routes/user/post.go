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
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

//Post func
func Post(w http.ResponseWriter, r *http.Request) {
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	res, err := client.CreateUser(context.Background(), &proto.CreateUserReq{
		User: &proto.User{
			ID:       int32(user.ID),
			Username: user.Username,
			FullName: user.FullName,
			Password: string(hashedPassword),
		}})
	if err != nil {
		httputil.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	httputil.WriteResponse(w, "User created successfully", map[string]string{
		"access_token":  res.Token,
		"refresh_token": res.RefreshToken,
		"user_id":       strconv.Itoa(int(res.UserID)),
	})
}
