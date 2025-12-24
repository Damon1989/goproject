package api

import "github.com/gin-gonic/gin"

type ApiController struct {
}

func (a ApiController) Index(c *gin.Context) {
	c.String(200, "我是一个api接口")
}

func (a ApiController) UserList(c *gin.Context) {
	c.String(200, "我是一个api接口 userlist")
}

func (a ApiController) Plist(c *gin.Context) {
	c.String(200, "我是一个api接口")
}

func (a ApiController) Cart(c *gin.Context) {
	c.String(200, "我是一个api接口")
}
