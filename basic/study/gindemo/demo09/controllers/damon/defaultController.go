package damon

import (
	"demo/models"
	"github.com/gin-gonic/gin"
	"time"
)

type DefaultController struct {
}

func (d DefaultController) Index(c *gin.Context) {

	t := models.UnixToTime(int(time.Now().Unix()))

	c.String(200, "首页:"+t)
}

func (d DefaultController) News(c *gin.Context) {
	c.String(200, "新闻")
}
