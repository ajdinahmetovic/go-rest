package main

import (
	"net/http"
	"time"

	"github.com/ajdinahmetovic/go-rest/routes"
	"github.com/ajdinahmetovic/item-service/logger"
)

func main() {
	logger.InitLogger()
	var router = routes.CreateRoutes()
	err := http.ListenAndServe(":8000", &router)
	if err != nil {
		return
	}
	logger.Info("Server started", "time", time.Now())
}
