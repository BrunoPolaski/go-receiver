package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Dia struct {
	Dia string `json:"dia"`
}

func Receive(c *gin.Context) {
	fmt.Println("Received request")

	var dia Dia

	if err := c.BindJSON(&dia); err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request",
		})
		return
	}

	fmt.Printf("Dia: %v\n", dia.Dia)

	fmt.Printf("Request: %v\n", c.Request)
	c.JSON(200, gin.H{
		"message": "received",
	})
}
