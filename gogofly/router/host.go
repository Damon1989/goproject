package router

import (
	"github.com/damon/gogofly/api"
	"github.com/gin-gonic/gin"
)

func InitHostRoutes() {
	RegisterRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		hostApi := api.NewHostApi()

		rgAuthHost := rgAuth.Group("host")
		{
			rgAuthHost.POST("/shutdown", hostApi.Shutdown)
		}
	})
}
