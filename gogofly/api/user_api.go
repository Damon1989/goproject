package api

import (
	"github.com/damon/gogofly/service"
	"github.com/damon/gogofly/service/dto"
	"github.com/damon/gogofly/utils"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() UserApi {
	return UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
	}
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
func (m UserApi) Login(c *gin.Context) {
	var iUserLoginDTO dto.UserLoginDTO

	if err := m.BuildRequest(BuildRequestOptions{Ctx: c, DTO: &iUserLoginDTO}).GetError(); err != nil {
		return
	}

	iUser, err := m.Service.Login(iUserLoginDTO)
	if err != nil {
		m.Fail(ResponseJson{
			Msg: err.Error(),
		})
		return
	}

	token, _ := utils.GenerateToken(iUser.ID, iUser.Name)

	m.OK(ResponseJson{
		Data: gin.H{
			"token": token,
			"user":  iUser,
		},
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
