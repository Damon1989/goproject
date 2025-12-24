package main

import (
	"demo02/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.AdminRouterInit(r)
	routers.ApiRoutersInit(r)
	routers.DefaultRoutersInit(r)

	r.Run()
}
