package router

import (
	"github.com/damon/gogofly/api"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes() {
	RegisterRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()
		rgPublicUser := rgPublic.Group("user").Use(func() gin.HandlerFunc {
			return func(c *gin.Context) {
				//c.AbortWithStatusJSON(http.StatusOK, gin.H{
				//	"msg": "login middleware",
				//})

			}
		}())
		{
			rgPublicUser.POST("/login", userApi.Login)
		}

		rgAuthUser := rgAuth.Group("user")
		{
			//rgAuthUser.GET("", func(c *gin.Context) {
			//	c.AbortWithStatusJSON(http.StatusOK, gin.H{
			//		"data": []map[string]any{
			//			{"id": 1, "name": "user1"},
			//			{"id": 2, "name": "user2"},
			//		},
			//	})
			//})
			//rgAuthUser.GET("/:id", func(c *gin.Context) {
			//	c.AbortWithStatusJSON(http.StatusOK, gin.H{
			//		"data": map[string]any{
			//			"id":   c.Param("id"),
			//			"name": "user1",
			//		},
			//	})
			//})
			rgAuthUser.POST("", userApi.AddUser)
			rgAuthUser.GET("/:id", userApi.GetUserById)

		}
	})
}
