package routes

import (
	"github.com/ajdinahmetovic/go-rest/db"
	"github.com/ajdinahmetovic/go-rest/routes/item"
	"github.com/ajdinahmetovic/go-rest/routes/user"
	"github.com/gorilla/mux"
)

//Router instance
var router = mux.NewRouter()

//CreateRoutes function initializes routes
func CreateRoutes() mux.Router {

	db.ConnectDB()

	//Item routes
	router.HandleFunc("/item", item.Put).Methods("PUT")
	router.HandleFunc("/item", item.GetItem).Methods("GET")
	router.HandleFunc("/item", item.Post).Methods("POST")
	router.HandleFunc("/item", item.Delete).Methods("DELETE")

	//User routes
	router.HandleFunc("/user", user.Post).Methods("POST")
	router.HandleFunc("/user", user.Get).Methods("GET")
	router.HandleFunc("/user", user.Delete).Methods("DELETE")
	router.HandleFunc("/user", user.Put).Methods("PUT")
	return *router
}

//GetRouter returns router
func GetRouter() mux.Router {
	return *router
}
