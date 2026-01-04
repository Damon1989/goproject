package dto

type UserLoginDTO struct {
	Name     string `json:"name" binding:"required,first_is_a" message:"用户名首字母必须是a" required_err:"用户名不能为空"`
	Password string `json:"password" binding:"required"`
}
