package main

import (
	"fmt"
	"github.com/buguzei/effective-mobile/auth/delivery/grpc"
	"github.com/buguzei/effective-mobile/auth/internal/repo/refresh/redis"
	"github.com/buguzei/effective-mobile/auth/internal/repo/user/postgres"
	"github.com/buguzei/effective-mobile/auth/internal/server"
	"github.com/buguzei/effective-mobile/auth/internal/usecase"
	"github.com/buguzei/effective-mobile/pkg/logging"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// init Logger
	var logger logging.Logger = logging.NewLogrus("debug")
	logger = logger.Named("auth.main")

	// reading .env file
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("reading env file failed", logging.Fields{
			"error": err,
		})
	}

	redisCli := redis.New()
	pgrsCli := postgres.New()

	uc := usecase.New(redisCli, pgrsCli)

	handler := grpc.New(uc)

	logger.Info(fmt.Sprintf("server started on port %s", os.Getenv("PORT")), logging.Fields{})
	if err = server.Run(handler); err != nil {
		panic(err)
	}
}
