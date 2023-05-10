package model

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"ink-web/src/global"
	"ink-web/src/util/log"
)

type User struct {
	gorm.Model
	UUID     string `json:"uuid" gorm:"unique"`
	Username string `json:"uername" des:"用户名"`
	Password string `json:"password" des:"密码"`
	Mobile   string `json:"mobile"`
}

func GetUserByName(username string) *[]User {
	var user []User
	err := global.Mysql.Unscoped().Model(&User{}).Where("username = ?", username).Find(&user).Error
	if err != nil {
		log.Error("GetUser", zap.Error(err))
		return nil
	}
	return &user
}
