package router

import (
	// "net/http"

	docs "chat-gin.com/docs"
	"chat-gin.com/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RouterInit() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/index", service.GetIndex)
	r.GET("/user/getUserlist", service.GetUserList)
	return r
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
