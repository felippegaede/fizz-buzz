// Package api contains all routes
package api

import "github.com/gin-gonic/gin"

type Server struct {
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer() *Server {
	router := gin.Default()

	server := &Server{}

	router.GET("/fizz-buzz/:number", server.fizzBuzz)

	server.router = router

	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(serverAddress string) error {
	return server.router.Run(serverAddress)
}
