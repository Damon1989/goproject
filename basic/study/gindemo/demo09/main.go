package main

import (
	"demo/models"
	"demo/routers"
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	r := gin.Default()

	// 自定义模板函数  注意要把这个函数放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})
	routers.AdminRouterInit(r)
	routers.ApiRoutersInit(r)
	routers.DefaultRoutersInit(r)

	r.Run()
}
