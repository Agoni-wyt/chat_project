package service

import (
	"net/http"

	"chat-gin.com/models"
	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context){
	data := models.GetUserList()
	c.JSON(http.StatusOK,gin.H{
		"data": data,
		"message": "User List",
		"code": "0",
	})
}