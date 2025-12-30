package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type IFnRegisterRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRoutes []IFnRegisterRoute
)

func RegisterRoute(fn IFnRegisterRoute) {
	if fn == nil {
		return
	}
	gfnRoutes = append(gfnRoutes, fn)
}

func InitBasePlatformRoutes() {
	InitUserRoutes()
}

func InitRouter() {
	r := gin.Default()
	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")

	InitBasePlatformRoutes()

	for _, fnRegisterRoute := range gfnRoutes {
		fnRegisterRoute(rgPublic, rgAuth)
	}
	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8999"
	}
	err := r.Run(fmt.Sprintf(":%s", stPort))
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
