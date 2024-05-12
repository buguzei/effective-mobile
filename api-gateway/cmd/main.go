package main

import (
	"fmt"
	delivery "github.com/buguzei/effective-mobile/api-gateway/delivery/http"
	_ "github.com/buguzei/effective-mobile/api-gateway/docs"
	"github.com/buguzei/effective-mobile/api-gateway/internal/server"
	"github.com/buguzei/effective-mobile/pkg/logging"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

// @title Car App API
// @version 1.0
// @description API Server For Car's Catalog

// @host localhost:8072
// @BasePath /

func main() {
	// init Logger
	var logger logging.Logger = logging.NewLogrus("debug")
	logger = logger.Named("api-gateway.main")

	// reading .env file
	err := godotenv.Load("./api-gateway/cmd/.env")
	if err != nil {
		logger.Fatal("reading env file failed", logging.Fields{
			"error": err,
		})
	}

	handler := delivery.NewHandler()
	s := server.NewServer(new(http.Server))

	// running server
	logger.Info(fmt.Sprintf("starting listen on port %s", os.Getenv("PORT")), logging.Fields{})
	err = s.Run(os.Getenv("PORT"), handler.InitRoutes())
	if err != nil {
		logger.Fatal("making migrations failed", logging.Fields{
			"error": err,
		})
	}
}
