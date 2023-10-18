package router

import (
	"github.com/gin-gonic/gin"
	"wxcloudrun-golang/controller"
)

func Router() {
	r := gin.Default()
	user := r.Group("/user")
	user.Use(controller.UserController)
	r.Run(":80")
}
