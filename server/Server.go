package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Init is used to initialize a server type struct
func (s *Server) Init() {
	if os.Getenv("PORT") == "" {
		s.Port = "8080"
	} else {
		s.Port = os.Getenv("PORT")
		log.Println("Got port " + s.Port)
		if s.Port == "" {
			log.Fatal("PORT must be Set")
		}
	}
	s.Router = gin.Default()
	s.Router.Use(cors.Default())
}

// AddGroup is used to add a RouterGroup to a server
func (s *Server) AddGroup(relativePath string) *SubGroup {
	G := SubGroup{}
	G.G = s.Router.Group(relativePath)
	G.Path = relativePath
	return &G
}

// AddGroupSub is used to add a RouterGroup to a SubGroup
func (s *SubGroup) AddGroupSub(relativePath string) *SubGroup {
	G := SubGroup{}
	G.G = s.G.Group(relativePath)
	G.Path = s.Path + relativePath
	return &G
}

// Start is used to Start a server
func (s Server) Start() {
	log.Fatal(s.Router.Run(fmt.Sprintf(":%s", s.Port)))
}
