package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/parmBody", func(c *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) //重新回写到缓存才能拿到parm1 parm2
		parm1 := c.PostForm("parm1")
		parm2 := c.PostForm("parm2")
		c.String(http.StatusOK, "%s %s %s", parm1, parm2, string(bodyBytes))
	})

	r.Run()
}
