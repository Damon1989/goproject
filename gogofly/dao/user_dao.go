package dao

import (
	"github.com/damon/gogofly/model"
)

var userDao *UserDao

type UserDao struct {
	BaseDao
}

func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{
			BaseDao: NewBaseDao(),
		}
	}
	return userDao
}

func (m *UserDao) GetUserByNameAndPassword(name string, password string) (user model.User) {
	var iUser model.User
	m.Orm.Where("name = ? and password = ?", name, password).First(&iUser)
	return iUser
}
