package main

import "github.com/gin-gonic/gin"

func main() {
	// Configuration
	gin.SetMode(gin.ReleaseMode)

	// Routing
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!")
	})

	// Running
	if err := r.Run(":80"); err != nil {
		panic(err)
	}
}
