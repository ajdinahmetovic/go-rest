package routes

import (
	"github.com/ajdinahmetovic/go-rest/routes/item"
	"github.com/gorilla/mux"
)

//Router instance
var router = mux.NewRouter()

//CreateRoutes function initializes routes
func CreateRoutes() mux.Router {
	router.HandleFunc("/item", item.Get).Methods("GET")
	router.HandleFunc("/item", item.Post).Methods("POST")
	return *router
}

//GetRouter returns router
func GetRouter() mux.Router {
	return *router
}
