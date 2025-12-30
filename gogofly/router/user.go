package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes() {
	RegisterRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		rgPublic.POST("/login", func(c *gin.Context) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg": "login success",
			})
		})
		rgAuthUser := rgAuth.Group("user")
		rgAuthUser.GET("", func(c *gin.Context) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"data": []map[string]any{
					{"id": 1, "name": "user1"},
					{"id": 2, "name": "user2"},
				},
			})
		})
		rgAuthUser.GET("/:id", func(c *gin.Context) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"data": map[string]any{
					"id":   c.Param("id"),
					"name": "user1",
				},
			})
		})
	})
}
