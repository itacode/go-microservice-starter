package main

import (
	"os"

	"github.com/itacode/go-microservice-starter/internal/env"
	"github.com/itacode/go-microservice-starter/internal/logger"
	"github.com/itacode/go-microservice-starter/internal/server"
)

func main() {
	env.LoadEnv()
	logger.InfoLog.Println("APP_ENV:", os.Getenv("APP_ENV"))
	logger.InfoLog.Println("APP_ADDR:", os.Getenv("APP_ADDR"))
	logger.InfoLog.Println("GIN_MODE:", os.Getenv("GIN_MODE"))
	server.Serve()
}
