package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/user/*action", func(c *gin.Context) { //设置获取name和id的参数
		c.String(200, "泛绑定～")
	})

	r.Run()
}
