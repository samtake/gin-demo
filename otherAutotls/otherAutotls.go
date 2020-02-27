package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "test")
	})
	autotls.Run(r, "www.XXX.ip")
}