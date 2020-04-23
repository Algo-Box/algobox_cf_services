package main

import (
	"github.com/algobox_cf_services/api"
	"github.com/algobox_cf_services/server"
)

func main() {
	server := server.Server{}
	server.Init()

	APIGroup := server.AddGroup("/api")
	api.Init(APIGroup)

	server.Start()
}
