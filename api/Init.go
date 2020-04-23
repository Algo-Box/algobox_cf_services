package api

import (
	"github.com/algobox_cf_services/server"
	"github.com/algobox_cf_services/user"
)

// Init Function initializes the api subgroup
func Init(g *server.SubGroup) {
	g.G.GET("/", handleHome)
	userRoute := g.AddGroupSub("/user")
	user.Init(userRoute)
}
