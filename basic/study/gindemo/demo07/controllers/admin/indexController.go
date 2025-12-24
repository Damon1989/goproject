package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
}

func (u IndexController) Index(c *gin.Context) {
	c.String(http.StatusOK, "首页列表")
}
