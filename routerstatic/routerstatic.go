package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 静态文件夹绑定+路由，有两种写法：
	r.Static("/assets", "./assets")
	r.StaticFS("/static", http.Dir("static"))

	//又或者路由+资源
	r.StaticFile("/favicon.ico", "/favicon.ico")
	r.Run()
}
