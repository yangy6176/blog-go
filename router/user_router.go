package router

import (
	"github.com/gin-gonic/gin"
	"wxcloudrun-golang/controller"
)

func UserRouter() {
	r := gin.Default()
	r.GET("/login", controller.UserController)
}
