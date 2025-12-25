package main

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func main() {
	gin.SetMode(gin.ReleaseMode)
	//1. 创建默认的路由引擎
	r := gin.Default()
	//2. 绑定路由规则，执行的函数
	r.GET("/index", Index)
	//3. 监听端口，默认8080
	r.Run("127.0.0.1:8080")

}
