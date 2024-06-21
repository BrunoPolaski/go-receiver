package main

import (
	"github.com/BrunoPolaski/go-receiver/controller/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	routes.InitRoute(&r.RouterGroup)

	if err := r.Run("localhost:8088"); err != nil {
		panic(err)
	}
}
