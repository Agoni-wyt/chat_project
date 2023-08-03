package models

import (
	"time"

	"chat-gin.com/utils"
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Password string

	Phone string `gorm:"unique"`
	Email string `gorm:"unique"`

	Identity   string
	ClientIp   string
	ClientPort string

	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time `gorm:"column:login_out_time" json:"login_out_time"`

	IsLogout   bool
	DeviveInfo string
}

// 实体类
func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList()[]*UserBasic{
	data := make([]*UserBasic,5)
	utils.DB.Find(&data)
	return data
}