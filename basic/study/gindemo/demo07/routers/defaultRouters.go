package routers

import (
	"demo07/controllers/damon"
	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", damon.DefaultController{}.Index)
		defaultRouters.GET("/news", damon.DefaultController{}.News)
	}
}
