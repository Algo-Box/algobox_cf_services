package user

import (
	"errors"

	"github.com/algobox_cf_services/structs"

	"github.com/algobox_cf_services/utils"
	"github.com/gin-gonic/gin"
)

func handleUser(c *gin.Context) {
	handle := c.Param("username")
	resp := getAllSubs(handle)
	if resp == nil {
		err := errors.New("InternalServerError")
		utils.WriteResponse(err, nil, c)
		return
	}
	utils.WriteResponse(nil, resp, c)
}

func handleHome(c *gin.Context) {
	req := structs.UserRequest{}
	if err := c.BindJSON(&req); err != nil {
		utils.WriteResponse(err, nil, c)
		return
	}
	utils.WriteResponse(nil, getDashboardData(req.CfHandle), c)
}
