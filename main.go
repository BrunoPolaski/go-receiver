package main

import (
	"fmt"
	"time"

	"github.com/BrunoPolaski/go-receiver/controller/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	fmt.Println("Starting application...")

	routes.InitRoute(&r.RouterGroup)
	go func() {
		chars := []rune{'|', '\\', '-', '/'}
		for {
			for _, char := range chars {
				fmt.Printf("\r%s", string(char))
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	fmt.Println("\nRunning!")

	if err := r.Run(":8088"); err != nil {
		panic(err)
	}
}
