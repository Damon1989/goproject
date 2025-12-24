package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "首页")
		})
		defaultRouters.GET("/news", func(c *gin.Context) {
			c.String(http.StatusOK, "新闻")
		})
	}
}

func main() {

}
