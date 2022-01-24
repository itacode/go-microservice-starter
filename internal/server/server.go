package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itacode/go-microservice-starter/internal/api/router"
)

func NewServer() *http.Server {
	fmt.Println("gin.Mode():", gin.Mode())

	addr := os.Getenv("APP_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	router := router.NewRouter()

	server := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()

	return server
}
