package routes

import (
	"github.com/ajdinahmetovic/go-rest/db"
	"github.com/ajdinahmetovic/go-rest/routes/item"
	"github.com/gorilla/mux"
)

//Router instance
var router = mux.NewRouter()

//CreateRoutes function initializes routes
func CreateRoutes() mux.Router {

	db.ConnectDB()

	router.HandleFunc("/item", item.GetItem).Methods("GET")
	router.HandleFunc("/item", item.Post).Methods("POST")
	return *router
}

//GetRouter returns router
func GetRouter() mux.Router {
	return *router
}
