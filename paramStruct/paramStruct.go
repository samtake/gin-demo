package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name     string    `form:"name"` //设置tag `form`可以由参数转变成结构体
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-03"`
}

func main() {
	r := gin.Default()
	r.GET("/test", testHandler)
	r.POST("/test", testHandler)
	r.Run()
}

func testHandler(c *gin.Context) {
	var person Person
	//根据请求content-type来作不同的binding操作
	if err := c.ShouldBind(&person); err == nil {
		c.String(200, "%v", person)
	} else {
		c.String(200, "person bind error", err)
	}

}
