package main

import (
	"context"
	"net/http"

	"github.com/ajdinahmetovic/go-rest/config"
	"github.com/ajdinahmetovic/go-rest/es"
	"github.com/ajdinahmetovic/go-rest/logger"
	"github.com/ajdinahmetovic/go-rest/routes"
)

func main() {

	//Init zap logger
	logger.InitLogger()

	//Load AppConfig
	config.Load()

	//Init elasticsearch
	err := es.Init(context.Background())
	if err != nil {
		logger.Error("Elasticsearch failed", "err", err)
		return
	}
	logger.Info("Server ready")
	var router = routes.CreateRoutes()
	err = http.ListenAndServe(":8000", &router)
	if err != nil {
		return
	}
}
