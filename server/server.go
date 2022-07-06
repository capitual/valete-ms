package server

import (
	"log"

	"github.com/capitual/valete_ms/config"
	"github.com/capitual/valete_ms/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   config.GetConfig().ServerPort,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run(":" + s.port))
}
