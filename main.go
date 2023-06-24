package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 处理GET请求
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	// 处理POST请求
	r.POST("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "world")
		_, err := fmt.Fprintln(c, "Hello,", name)
		if err != nil {
			return
		}
	})

	// 启动服务器
	r.Run(":8080")
}
