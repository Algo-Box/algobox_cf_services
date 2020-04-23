package server

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Init is used to initialize a server type struct
func (s *Server) Init(port ...string) {
	if len(port) >= 1 {
		s.Port = port[0]
	} else {
		s.Port = "8080"
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
