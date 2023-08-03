package models

import "github.com/jinzhu/gorm"

type UserBasic struct {
	gorm.Model
	Name       string
	Password   string

	Phone      string
	Email      string

	Identity   string
	ClientIp   string
	ClientPort string

	LoginTime uint64
	HeartbeatTime uint64
	LogOutTime uint64

	IsLogout bool
	DeviveInfo string
}
// 实体类
func(table *UserBasic)TableName() string {
	return "user_basic"
}
