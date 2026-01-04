package service

import (
	"errors"

	"github.com/damon/gogofly/dao"
	"github.com/damon/gogofly/model"
	"github.com/damon/gogofly/service/dto"
)

var userService *UserService

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			Dao: dao.NewUserDao(),
		}
	}
	return userService
}

func (m *UserService) Login(dto dto.UserLoginDTO) (model.User, error) {
	var errResult error
	iUser := m.Dao.GetUserByNameAndPassword(dto.Name, dto.Password)
	if iUser.ID == 0 {
		errResult = errors.New("用户名或密码错误")
	}
	return iUser, errResult
}
