package api

import (
	"github.com/damon/gogofly/service/dto"
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
	var iUserLoginDTO dto.UserLoginDTO
	if err := ctx.ShouldBind(&iUserLoginDTO); err != nil {
		Fail(ctx, ResponseJson{
			Msg: err.Error(),
		})
		return
	}
	OK(ctx, ResponseJson{
		Data: iUserLoginDTO,
	})
	//ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
	//	"msg": "login success",
	//})

	//OK(ctx, ResponseJson{
	//	Msg: "login success",
	//})

	//Fail(ctx, ResponseJson{
	//	Code: 9001,
	//	Msg:  "login fail",
	//})

}
