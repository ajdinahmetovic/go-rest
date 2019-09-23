package main

import (
	"net/http"
	"time"

	"github.com/ajdinahmetovic/go-rest/routes"
	"github.com/ajdinahmetovic/item-service/logger"
)

func main() {
	var router = routes.CreateRoutes()
	err := http.ListenAndServe(":3000", &router)
	if err != nil {
		logger.Error("Failed to start server", "time", time.Now(), "err", err)
		return
	}
	logger.Info("Server started", "time", time.Now())

}
