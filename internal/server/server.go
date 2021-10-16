package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func NewServer() *http.Server {
	fmt.Println("gin.Mode()", gin.Mode())

	addr := os.Getenv("APP_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

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
