package main

import (
	"fmt"
	"net/http"

	"github.com/ajdinahmetovic/go-rest/routes"
)

func main() {
	fmt.Println("Server starting")
	var router = routes.CreateRoutes()
	err := http.ListenAndServe(":3000", &router)
	if err != nil {
		fmt.Println("Server failed ", err)
	}
	fmt.Println("Server started")
}
