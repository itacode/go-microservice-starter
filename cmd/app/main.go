package main

import (
	"fmt"
	"os"

	"github.com/itacode/go-microservice-starter/internal/env"
	"github.com/itacode/go-microservice-starter/internal/server"
)

func main() {
	env.LoadEnv()
	fmt.Println("APP_ENV:", os.Getenv("APP_ENV"))
	fmt.Println("APP_ADDR:", os.Getenv("APP_ADDR"))
	fmt.Println("GIN_MODE:", os.Getenv("GIN_MODE"))
	server.NewServer()
}
