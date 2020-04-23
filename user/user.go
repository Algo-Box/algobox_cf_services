package user

import (
	"github.com/algobox_cf_services/server"
)

const (
	// MaxSubs is the assumned maximum number of submissions a use can have
	MaxSubs = 10000
)

// Init is Used to Initialize the user Instance
func Init(g *server.SubGroup) {
	g.G.GET("/:username", handleUser)
	g.G.GET("/", handleHome)
}
