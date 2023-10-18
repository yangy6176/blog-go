package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wxcloudrun-golang/module"
	"wxcloudrun-golang/util"
)

func UserLoginService(c *gin.Context) {
	var codeData module.WXLoginByCodeType
	if err := c.ShouldBind(&codeData); err != nil {
		c.JSON(http.StatusInternalServerError, util.DefaultResponse{
			Data:    struct{}{},
			Message: "参数缺失，请联系管理员",
			Success: false,
			Total:   0,
		})
	}

}
