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
	fmt.Printf("Request content length: %v\n", c.Request.ContentLength)
	fmt.Printf("Request content type: %v\n", c.Request.Header.Get("Content-Type"))
	fmt.Printf("Request user agent: %v\n", c.Request.UserAgent())
	fmt.Printf("Request protocol: %v\n", c.Request.Proto)
	fmt.Printf("Request authentication: %v\n", c.Request.Header.Get("Authorization"))
	val, _ := c.GetRawData()
	fmt.Printf("Request body: %v\n\n", string(val))
	fmt.Printf("---- End of request ----\n\n")
	c.JSON(200, gin.H{
		"message": "received",
	})
}
