package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ajdinahmetovic/go-rest/httputil"
	"github.com/elastic/go-elasticsearch"
)

//Get func
func Get(w http.ResponseWriter, r *http.Request) {
	httputil.EnableCors(&w, r)
	v := r.URL.Query()
	id, err := strconv.Atoi(v.Get("id"))
	if err != nil {
		id = 0
	}
	username := v.Get("username")
	fullname := v.Get("fullname")

	es, _ := elasticsearch.NewDefaultClient()
	log.Println(es.Info())

	//Search with elastcsearch

	//Search in postgress
	/*
		conn, err := grpc.Dial("service:4040", grpc.WithInsecure())
		if err != nil {
			httputil.WriteError(w, err, http.StatusInternalServerError)
			return
		}
			client := proto.NewUserServiceClient(conn)
			res, err := client.GetUser(context.Background(), &proto.GetUserReq{
				User: &proto.User{
					ID:       int32(id),
					Username: username,
					FullName: fullname,
				}})
			if err != nil {
				httputil.WriteError(w, err, http.StatusInternalServerError)
				return
			}
			httputil.WriteResponse(w, res.Message, res.Users)
	*/
}
