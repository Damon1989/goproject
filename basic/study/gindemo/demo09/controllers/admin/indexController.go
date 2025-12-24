package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
}

func (u IndexController) Index(c *gin.Context) {
	userName, _ := c.Get("userName")
	fmt.Println(userName)
	v, ok := userName.(string)
	if !ok {
		c.String(http.StatusOK, "首页列表")
	}
	c.String(http.StatusOK, "首页列表:"+v)
}
