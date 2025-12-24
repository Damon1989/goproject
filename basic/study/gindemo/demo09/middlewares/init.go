package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Index(c *gin.Context) {
	//判断用户是否登录

	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)

	c.Set("userName", "damon")

	// 定义一个goroutine统计日志
	cCopy := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done! in path: ", cCopy.Request.URL.Path)
	}()
}
