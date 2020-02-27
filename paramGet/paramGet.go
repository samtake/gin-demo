package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/testGetParam", func(c *gin.Context) { //设置获取name和id的参数
		param1 := c.Query("parm1")
		param2 := c.Query("parm2")

		c.String(http.StatusOK, "%s, %s", param1, param2)
	})

	r.Run()
}





