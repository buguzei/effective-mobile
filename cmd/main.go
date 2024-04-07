package main

import (
	"fmt"
	_ "github.com/buguzei/effective-mobile/docs"
	delivery "github.com/buguzei/effective-mobile/internal/delivery/http"
	"github.com/buguzei/effective-mobile/internal/repo/postgres"
	"github.com/buguzei/effective-mobile/internal/server"
	"github.com/buguzei/effective-mobile/internal/usecase"
	"github.com/buguzei/effective-mobile/pkg/logging"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"net/http"
	"os"
)

// @title Car App API
// @version 1.0
// @description API Server For Car's Catalog

// @host localhost:8087
// @BasePath /

func main() {
	// init Logger
	var logger logging.Logger = logging.NewLogrus("debug")
	logger = logger.Named("main")

	// reading .env file
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("reading env file failed", logging.Fields{
			"error": err,
		})
	}

	// initializing repository, use case, handler and server
	repo, err := postgres.NewPostgres(os.Getenv("DBCONN"))
	if err != nil {
		logger.Fatal("error of initializing postgres", logging.Fields{
			"error": err,
		})
	}

	uc := usecase.NewUseCase(repo)

	handler := delivery.NewHandler(uc)

	s := server.NewServer(new(http.Server))

	// making migrations
	if err = goose.SetDialect("postgres"); err != nil {
		logger.Fatal("setting dialect failed", logging.Fields{
			"error": err,
		})
	}

	err = goose.Up(repo.DB, "./migrations")
	if err != nil {
		logger.Fatal("making migrations failed", logging.Fields{
			"error": err,
		})
	}

	// running server
	logger.Info(fmt.Sprintf("starting listen on port %s", os.Getenv("PORT")), logging.Fields{})

	err = s.Run(os.Getenv("PORT"), handler.InitRoutes())
	if err != nil {
		logger.Fatal("making migrations failed", logging.Fields{
			"error": err,
		})
	}
}
