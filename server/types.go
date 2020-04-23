package server

import "github.com/gin-gonic/gin"

type (
	// Server is the server struct
	Server struct {
		Router *gin.Engine
		Port   string
	}

	// SubGroup is Subgroup type struct
	SubGroup struct {
		G    *gin.RouterGroup
		Path string
	}
)
