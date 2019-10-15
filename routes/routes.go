package routes

import (
	"github.com/ajdinahmetovic/go-rest/routes/item"
	"github.com/ajdinahmetovic/go-rest/routes/user"
	"github.com/gorilla/mux"
)

//Router instance
var router = mux.NewRouter()

//CreateRoutes function initializes routes
func CreateRoutes() mux.Router {
	//Item routes
	router.HandleFunc("/item", VerifyTokenMiddleware(item.Put)).Methods("PUT")
	router.HandleFunc("/item", VerifyTokenMiddleware(item.GetItem)).Methods("GET", "OPTIONS")
	router.HandleFunc("/item", VerifyTokenMiddleware(item.Post)).Methods("POST")
	router.HandleFunc("/item", VerifyTokenMiddleware(item.Delete)).Methods("DELETE")

	router.HandleFunc("/refresh", RefreshToken).Methods("POST")

	//User routes
	router.HandleFunc("/login", user.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/user", user.Post).Methods("POST", "OPTIONS")
	router.HandleFunc("/user", VerifyTokenMiddleware(user.Get)).Methods("GET")
	router.HandleFunc("/user", VerifyTokenMiddleware(user.Delete)).Methods("DELETE")
	router.HandleFunc("/user", VerifyTokenMiddleware(user.Put)).Methods("PUT")

	return *router
}

//GetRouter returns router
func GetRouter() mux.Router {
	return *router
}
