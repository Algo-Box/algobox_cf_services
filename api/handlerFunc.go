package api

import (
	"github.com/algobox_cf_services/utils"
	"github.com/gin-gonic/gin"
)

func handleHome(c *gin.Context) {
	r := Request{}
	x := c.BindJSON(&r)
	if x != nil {
		utils.WriteResponse(x, nil, c)
		return
	}
	utils.WriteResponse(nil, r, c)
}
