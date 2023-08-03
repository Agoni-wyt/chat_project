package service

import (
	"net/http"

	"chat-gin.com/models"
	"github.com/gin-gonic/gin"
)
// GetUserList
// @Tags 用户列表
// @Success 200 {json} json{"data": [],"message":"User List","code":"0"}
// @Router /user/getUserlist [get]
func GetUserList(c *gin.Context){
	data := models.GetUserList()
	c.JSON(http.StatusOK,gin.H{
		"data": data,
		"message": "User List",
		"code": "0",
	})
}