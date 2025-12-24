package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 配置路由
	r.GET("/", func(c *gin.Context) {
		c.String(200, "值：%v", "你好gin")
	})

	r.GET("/news", func(c *gin.Context) {
		c.String(200, "我是新闻界面!!!")
	})

	r.POST("/add", func(c *gin.Context) {
		c.String(http.StatusOK, "这是一个post请求 用于增加数据")
	})

	r.PUT("/edit", func(c *gin.Context) {
		c.String(200, "这是一个put请求 用于编辑数据")
	})

	r.DELETE("/del", func(c *gin.Context) {
		c.String(200, "这是一个delete请求 用于删除数据")
	})

	// r.Run() 启动HTTP服务，默认在0.0.0.0:8080端口 启动服务
	r.Run(":8080") // 启动一个web服务
}
