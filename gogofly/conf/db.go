package conf

import (
	"time"

	"github.com/damon/gogofly/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB() (*gorm.DB, error) {
	logMode := logger.Info
	if !viper.GetBool("mode.develop") {
		logMode = logger.Error
	}
	db, err := gorm.Open(mysql.Open(viper.GetString("db.dsn")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_", // 表名前缀，`User` 表为 `t_users`
			SingularTable: true,   // 使用单数表名，启用该选项后，`User` 表为 `user`
		},
		Logger: logger.Default.LogMode(logMode),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(viper.GetInt("MaxIdleConns"))
	sqlDB.SetMaxOpenConns(viper.GetInt("MaxOpenConns"))
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&model.User{})
	return db, nil

}
