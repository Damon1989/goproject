package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/damon/gogofly/dao"
	"github.com/damon/gogofly/global"
	"github.com/damon/gogofly/global/constants"
	"github.com/damon/gogofly/model"
	"github.com/damon/gogofly/service/dto"
	"github.com/damon/gogofly/utils"
	"github.com/spf13/viper"
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

func SetLoginUserTokenToRedis(uid uint, token string) error {
	return global.RedisClient.Set(strings.Replace(constants.LOGIN_USER_TOKEN_REDIS_KEY, "{ID}", strconv.Itoa(int(uid)), -1), token, viper.GetDuration("jwt.tokenExpire")*time.Minute)
}

func (m *UserService) Login(dto dto.UserLoginDTO) (model.User, string, error) {
	var errResult error
	var token string
	iUser, err := m.Dao.GetUserByName(dto.Name)
	if err != nil || !utils.CheckPassword(dto.Password, iUser.Password) {
		errResult = errors.New("用户名或密码错误")
	} else { // 登录成功 生成 token
		token, err = utils.GenerateToken(iUser.ID, iUser.Name)
		if err != nil {
			errResult = errors.New(fmt.Sprintf("generate token error: %v", err))
		}
	}
	return iUser, token, errResult
}

func (m *UserService) AddUser(iUserAddDTO *dto.UserAddDTO) error {
	if m.Dao.CheckUserNameExist(iUserAddDTO.Name) {
		return errors.New("username already exists")
	}
	return m.Dao.AddUser(iUserAddDTO)
}

func (m *UserService) GetUserById(iCommonIDDTO *dto.CommonIDDTO) (model.User, error) {
	return m.Dao.GetUserById(iCommonIDDTO.ID)
}

func (m *UserService) GetUserList(iUserListDTO *dto.UserListDTO) ([]model.User, int64, error) {
	return m.Dao.GetUserList(iUserListDTO)
}

func (m *UserService) UpdateUser(iUserUpdateDTO *dto.UserUpdateDTO) error {
	if iUserUpdateDTO.ID == 0 {
		return errors.New("id can not be empty")
	}

	return m.Dao.UpdateUser(iUserUpdateDTO)
}

func (m *UserService) DeleteUserById(iCommonIDDTO *dto.CommonIDDTO) error {
	if iCommonIDDTO.ID == 0 {
		return errors.New("id can not be empty")
	}
	return m.Dao.DeleteUserById(iCommonIDDTO.ID)
}
