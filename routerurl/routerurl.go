package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/:name/:id", func(c *gin.Context) { //设置获取name和id的参数
		c.JSON(200, gin.H{
			"name": c.Param("name"),
			"id":   c.Param("id"),
		})
	})

	r.Run()
}
