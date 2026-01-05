package dto

import (
	"github.com/damon/gogofly/model"
)

type UserLoginDTO struct {
	Name     string `json:"name" binding:"required,first_is_a" message:"用户名首字母必须是a" required_err:"用户名不能为空"`
	Password string `json:"password" binding:"required"`
}

// 添加用户相关的 DTO
type UserAddDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" form:"name" binding:"required" message:"用户名不能为空"`
	RealName string `json:"real_name" form:"real_name" binding:"required"`
	Avatar   string `json:"avatar" form:"avatar" binding:"required"`
	Mobile   string `json:"mobile" form:"mobile" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password,omitempty" binding:"required" message:"密码不能为空"`
}

func (m *UserAddDTO) ConvertToModel(iUser *model.User) {
	iUser.Name = m.Name
	iUser.RealName = m.RealName
	iUser.Avatar = m.Avatar
	iUser.Mobile = m.Mobile
	iUser.Email = m.Email
	iUser.Password = m.Password
}
