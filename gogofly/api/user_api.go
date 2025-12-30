package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

func NewUserApi() UserApi {
	return UserApi{}
}

// @Summary 用户登录
// @Description 用户登录接口
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/public/user/login [post]
func (u UserApi) Login(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg": "login success",
	})
}
