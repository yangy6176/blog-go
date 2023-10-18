package controller

import (
	"github.com/gin-gonic/gin"
	"wxcloudrun-golang/service"
)

func UserController(r *gin.Engine) {
	user := r.Group("/user")
	user.GET("/login", service.UserLoginService)
}
