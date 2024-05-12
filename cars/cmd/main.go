package main

import (
	"fmt"
	"github.com/buguzei/effective-mobile/cars/internal/delivery/grpc"
	"github.com/buguzei/effective-mobile/cars/internal/repo/postgres"
	"github.com/buguzei/effective-mobile/cars/internal/server"
	"github.com/buguzei/effective-mobile/cars/internal/usecase"
	"github.com/buguzei/effective-mobile/pkg/logging"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// init Logger
	var logger logging.Logger = logging.NewLogrus("debug")
	logger = logger.Named("main")

	// reading .env file
	err := godotenv.Load("./cars/cmd/.env")
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

	defer func() {
		err = repo.DB.Close()
		if err != nil {
			logger.Error("error of closing db", logging.Fields{
				"error": err,
			})
		}
	}()

	uc := usecase.NewUseCase(repo)

	handler := grpc.New(uc)

	// making migrations
	//if err = goose.SetDialect("postgres"); err != nil {
	//	logger.Fatal("setting dialect failed", logging.Fields{
	//		"error": err,
	//	})
	//}
	//
	//err = goose.Up(repo.DB, "./migrations")
	//if err != nil {
	//	logger.Fatal("making migrations failed", logging.Fields{
	//		"error": err,
	//	})
	//}

	// running server
	logger.Info(fmt.Sprintf("starting listen on port %s", os.Getenv("PORT")), logging.Fields{})

	err = server.Run(handler)
	if err != nil {
		panic(err)
	}
}
