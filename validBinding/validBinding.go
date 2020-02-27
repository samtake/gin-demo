package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Age     int    `form:"Age"  binding:"required,gt=10"`
	Name    string `form:"Name" binding:"required"`
	Address string `form:"Address" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/test", testHandler)
	r.Run()
}

func testHandler(c *gin.Context) {
	var person Person
	//根据请求content-type来作不同的binding操作
	if err := c.ShouldBind(&person); err != nil {
		c.String(500, "%v", err)
		c.Abort()
		return
	}
	c.String(200, "%v", person)

}
