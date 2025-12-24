package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func initMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("1-我是一个中间件")
	// 调用该请求的剩余处理程序
	c.Next()

	fmt.Println("2-我是一个中间件")

	end := time.Now().UnixNano()
	fmt.Println("请求时间：", end-start)
}

func main() {
	r := gin.Default()

	//routers.AdminRouterInit(r)
	//routers.ApiRoutersInit(r)
	//routers.DefaultRoutersInit(r)

	// 全局中间件
	r.Use(initMiddleware)

	r.GET("/", func(c *gin.Context) {
		fmt.Println("这是一个首页")
		time.Sleep(1 * time.Second)
		c.String(200, "gin首页")
	})

	r.GET("/news", func(c *gin.Context) {
		c.String(200, "新闻首页")
	})

	r.GET("/login", initMiddleware, func(c *gin.Context) {
		c.String(200, "login")
	})

	r.Run()
}
