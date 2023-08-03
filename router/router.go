package router

import (
	// "net/http"

	"chat-gin.com/service"
	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine{
	r := gin.Default()
	r.GET("/index",service.GetIndex)
	r.GET("/user/getUserlist",service.GetUserList)
	return r
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
