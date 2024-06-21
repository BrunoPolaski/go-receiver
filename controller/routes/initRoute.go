package routes

import (
	"github.com/BrunoPolaski/go-receiver/controller"
	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.RouterGroup) {
	r.POST("/receive", controller.Receive)
}
