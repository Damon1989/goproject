package router

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/damon/gogofly/docs"
	"github.com/damon/gogofly/global"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

func initBasePlatformRoutes() {
	InitUserRoutes()
}

func registerCustomValidator() {
	global.Logger.Info("register custom validator")
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		fmt.Println("register custom validator")
		validate.RegisterValidation("first_is_a", func(fl validator.FieldLevel) bool {
			fmt.Println("2")
			if value, ok := fl.Field().Interface().(string); ok {
				if value != "" && value[0] == 'a' {
					return true
				}
			}
			return false
		})
	}
}

func InitRouter() {
	// 使用 signal.NotifyContext 创建一个会在接收到 SIGINT 或 SIGTERM 时取消的 Context
	// 这样我们可以在接收到退出信号时触发优雅关闭流程
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	// 创建 Gin 默认引擎并定义两个路由组：公共与鉴权
	r := gin.Default()
	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")

	// 初始化基础平台路由（例如用户路由）
	initBasePlatformRoutes()

	// 注册自定义验证器
	registerCustomValidator()

	// 调用所有已注册的路由注册函数，将路由组传入以完成具体路由的挂载
	for _, fnRegisterRoute := range gfnRoutes {
		fnRegisterRoute(rgPublic, rgAuth)
	}

	// 集成swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 从配置读取端口；若未设置则使用默认端口 8999
	stPort := viper.GetString("server.port")
	if stPort == "" {
		stPort = "8999"
	}

	// 创建并配置 HTTP Server，使用 Gin 引擎作为处理器
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", stPort),
		Handler: r,
	}

	// 在独立的 goroutine 中启动服务器，避免阻塞后续的信号监听
	go func() {
		// 增加启动日志 测试日志组件性能
		//for i := 0; i < 100000; i++ {
		//	global.Logger.Info(fmt.Sprintf("Starting server on port %s", stPort))
		//}
		global.Logger.Info(fmt.Sprintf("Starting server on port %s", stPort))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// TODO: 记录日志
			global.Logger.Error(fmt.Sprintf("Failed to start server: %v", err))
			//fmt.Println(fmt.Sprintf("Failed to start server: %v", err))
			return
		}
	}()

	// 等待接收到终止信号（由上面的 signal.NotifyContext 控制）
	<-ctx.Done()

	// 收到退出信号后，创建一个带超时的 Context 用于优雅关闭（最多等待 5 秒）
	ctx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()
	if err := server.Shutdown(ctx); err != nil {
		// TODO: 记录日志
		global.Logger.Error(fmt.Sprintf("Failed to shutdown server: %v", err))
		//fmt.Println(fmt.Sprintf("Failed to shutdown server: %v", err))
	}
	global.Logger.Info("stop server gracefully")
}
