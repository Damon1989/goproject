package api

import (
	"github.com/damon/gogofly/service"
	"github.com/damon/gogofly/service/dto"
	"github.com/damon/gogofly/utils"
	"github.com/gin-gonic/gin"
)

const (
	ERR_CODE_ADD_USER       = 10011
	ERR_CODE_GET_USER_BY_ID = 10012
	ERR_CODE_GET_USER_LIST  = 10013
	ERR_CODE_UPDATE_USER    = 10014
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

// @Summary 添加用户
// @Description 添加用户接口
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param name formData string true "用户名"
// @Param real_name formData string true "真实姓名"
// @Param avatar formData string true "头像"
// @Param mobile formData string true "手机号"
// @Param email formData string true "邮箱"
// @Param password formData string true "密码"
// @Success 200 {object} map[string]interface{}
func (m UserApi) AddUser(c *gin.Context) {
	var iUserAddDTO dto.UserAddDTO

	if err := m.BuildRequest(BuildRequestOptions{Ctx: c, DTO: &iUserAddDTO}).GetError(); err != nil {
		return
	}

	if err := m.Service.AddUser(&iUserAddDTO); err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_ADD_USER,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Msg:  "add user success",
		Data: iUserAddDTO,
	})
}

func (m UserApi) GetUserById(c *gin.Context) {
	var iCommonIDDTO dto.CommonIDDTO

	if err := m.BuildRequest(BuildRequestOptions{Ctx: c, DTO: &iCommonIDDTO, BindUri: true}).GetError(); err != nil {
		return
	}

	iUser, err := m.Service.GetUserById(&iCommonIDDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_BY_ID,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseJson{
		Data: iUser,
	})
}

func (m UserApi) GetUserList(c *gin.Context) {
	var iUserListDTO dto.UserListDTO

	if err := m.BuildRequest(BuildRequestOptions{Ctx: c, DTO: &iUserListDTO}).GetError(); err != nil {
		return
	}

	iUserList, total, err := m.Service.GetUserList(&iUserListDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_GET_USER_LIST,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseJson{
		Data:  iUserList,
		Total: total,
	})
}

func (m UserApi) UpdateUser(c *gin.Context) {
	var iUserUpdateDTO dto.UserUpdateDTO
	//strId := c.Param("id")
	//fmt.Println("strId:", strId)
	//id, _ := strconv.Atoi(strId)
	//uid := uint(id)
	//iUserUpdateDTO.ID = uid

	if err := m.BuildRequest(BuildRequestOptions{Ctx: c, DTO: &iUserUpdateDTO, BindAll: true}).GetError(); err != nil {
		return
	}

	err := m.Service.UpdateUser(&iUserUpdateDTO)
	if err != nil {
		m.ServerFail(ResponseJson{
			Code: ERR_CODE_UPDATE_USER,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseJson{})

}
