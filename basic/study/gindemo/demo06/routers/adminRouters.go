package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminRouterInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "后台首页")
		})

		adminRouters.GET("/user", func(c *gin.Context) {
			c.String(http.StatusOK, "用户列表")
		})

		adminRouters.GET("/user/add", func(c *gin.Context) {
			c.String(http.StatusOK, "用户列表-add")
		})

		adminRouters.GET("/user/edit", func(c *gin.Context) {
			c.String(http.StatusOK, "用户列表-edit")
		})

		adminRouters.GET("/article", func(c *gin.Context) {
			c.String(http.StatusOK, "新闻列表")
		})

		adminRouters.GET("/article/add", func(c *gin.Context) {
			c.String(http.StatusOK, "新闻列表-add")
		})

		adminRouters.GET("/article/edit", func(c *gin.Context) {
			c.String(http.StatusOK, "新闻列表-edit")
		})
	}
}

func main() {

}
