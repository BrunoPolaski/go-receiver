package routes

import (
	"github.com/BrunoPolaski/go-receiver/controller"
	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.RouterGroup) {
	r.POST("/receive", controller.Receive)
	r.GET("/receive", controller.Receive)
	r.PUT("/receive", controller.Receive)
	r.DELETE("/receive", controller.Receive)
	r.PATCH("/receive", controller.Receive)
	r.OPTIONS("/receive", controller.Receive)
	r.HEAD("/receive", controller.Receive)
}
