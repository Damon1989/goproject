package model

import (
	"github.com/damon/gogofly/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	RealName string `json:"real_name" gorm:"type:varchar(100);not null"`
	Avatar   string `json:"avatar" gorm:"type:varchar(100);not null"`
	Mobile   string `json:"mobile" gorm:"type:varchar(100);not null"`
	Email    string `json:"email" gorm:"type:varchar(100);not null"`
	Password string `json:"-" gorm:"type:varchar(100);not null"`
}

func (u *User) Encrypt() error {
	stHash, err := utils.Encrypt(u.Password)
	if err == nil {
		u.Password = stHash
	}
	return err
}

func (u *User) BeforeCreate(orm *gorm.DB) error {
	return u.Encrypt()
}
