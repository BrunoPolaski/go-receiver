package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Receive(c *gin.Context) {
	fmt.Printf("\n\n---- Received request ----\n\n")
	fmt.Printf("Request method: %v\n", c.Request.Method)
	fmt.Printf("Request host: %v\n", c.Request.Host)
	fmt.Printf("Request URL: %v\n", c.Request.URL)
	fmt.Printf("Request remote address: %v\n", c.Request.RemoteAddr)
	fmt.Printf("Request protocol: %v\n", c.Request.Proto)
	fmt.Println("Request headers:")
	for k, v := range c.Request.Header {
		fmt.Printf("  %v: %v\n", k, v)
	}
	if val, _ := c.GetRawData(); len(val) > 0 {
		fmt.Printf("Request body: %v", string(val))
	}
	fmt.Printf("\n---- End of request ----\n\n")
	c.JSON(200, gin.H{
		"message": "received",
	})
}
