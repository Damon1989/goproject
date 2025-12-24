package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleController struct {
}

func (a ArticleController) Index(c *gin.Context) {
	c.String(http.StatusOK, "文章列表")
}

func (a ArticleController) Add(c *gin.Context) {
	c.String(http.StatusOK, "文章列表-add")
}

func (a ArticleController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "文章列表-edit")
}
