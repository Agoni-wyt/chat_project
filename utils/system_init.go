package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB

func InitConfig(){
	viper.SetConfigName("conf")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

}

func InitMySQL()(err error){
	// fmt.Println(viper.GetString("mysql.dn"))
	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}