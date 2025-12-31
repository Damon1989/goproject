package cmd

import (
	"fmt"

	"github.com/damon/gogofly/conf"
	"github.com/damon/gogofly/global"
	"github.com/damon/gogofly/router"
	"github.com/damon/gogofly/utils"
)

func Start() {
	var initErr error
	// 初始化配置：读取配置文件（例如 settings.yml）并设置全局配置
	conf.InitConfig()

	// 初始化日志：根据配置初始化日志器（日志级别、输出位置等）
	global.Logger = conf.InitLogger()

	db, err := conf.InitDB()
	if err != nil {
		initErr = utils.AppendError(initErr, err)
	}
	global.DB = db

	// 启动路由并开始 HTTP 服务：此调用会启动 Gin 服务器并阻塞，直到接收到退出信号并完成优雅关闭
	router.InitRouter()
}

func Clean() {
	fmt.Println("======================Clean======================")
}
