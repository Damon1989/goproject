package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

func (b BaseController) success(c *gin.Context) {
	c.String(http.StatusOK, "success")
}

func (b BaseController) error(c *gin.Context) {
	c.String(http.StatusOK, "error")
}
