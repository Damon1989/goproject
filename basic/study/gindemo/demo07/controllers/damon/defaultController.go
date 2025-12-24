package damon

import "github.com/gin-gonic/gin"

type DefaultController struct {
}

func (d DefaultController) Index(c *gin.Context) {
	c.String(200, "首页")
}

func (d DefaultController) News(c *gin.Context) {
	c.String(200, "新闻")
}
