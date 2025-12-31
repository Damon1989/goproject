package cmd

import (
	"fmt"

	"github.com/damon/gogofly/conf"
	"github.com/damon/gogofly/global"
	"github.com/damon/gogofly/router"
	"github.com/damon/gogofly/utils"
)

// Package cmd 提供应用程序的启动与清理入口逻辑。
// Start 会完成必要的初始化（配置、日志、数据库、Redis），
// 并启动路由（HTTP 服务）。

// Start 初始化应用所需的各种资源并启动路由服务器。
// 主要步骤：
// 1. 初始化配置（读取 settings.yml 等）
// 2. 初始化日志并赋值到全局变量
// 3. 初始化数据库连接并赋值到全局变量
// 4. 初始化 Redis 客户端并赋值到全局变量
// 5. 如果初始化过程中出现错误，合并并记录，然后 panic
// 6. 启动路由（阻塞，直到服务关闭）
func Start() {
	var initErr error
	// 初始化配置：读取配置文件（例如 settings.yml）并设置全局配置
	conf.InitConfig()

	// 初始化日志：根据配置初始化日志器（日志级别、输出位置等）
	// 并保存到全局变量，方便应用其余部分使用
	global.Logger = conf.InitLogger()

	// 初始化数据库连接，返回 db 对象和可能的错误
	db, err := conf.InitDB()
	// 将 db 赋值到全局变量，供其他模块使用
	global.DB = db
	if err != nil {
		// 将当前错误合并到 initErr 中，utils.AppendError 支持将多个错误串联
		initErr = utils.AppendError(initErr, err)
	}

	// 初始化 Redis 客户端
	rdClient, err := conf.InitRedis()
	// 保存到全局变量，供缓存/会话等使用
	global.RedisClient = rdClient
	if err != nil {
		// 合并错误，延迟统一处理
		initErr = utils.AppendError(initErr, err)
	}
	_ = global.RedisClient.Set("username", "gogofly")
	global.Logger.Info(global.RedisClient.Get("username"))

	// 如果存在任意初始化错误，记录并终止程序
	if initErr != nil {
		// 如果 Logger 已经可用，先记录错误详情
		if global.Logger != nil {
			global.Logger.Error(initErr.Error())
		}
		// 然后 panic 以便外层（或进程管理工具）感知到启动失败
		panic(initErr.Error())
	}

	// 启动路由并开始 HTTP 服务：此调用会启动 Gin 服务器并阻塞，
	// 直到接收到退出信号并完成优雅关闭
	router.InitRouter()
}

// Clean 提供一个简单的清理或占位函数，当前仅打印标记信息。
// 可在后续扩展中加入连接关闭、临时文件清理等逻辑。
func Clean() {
	fmt.Println("======================Clean======================")
}
