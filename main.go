package main

import (
	"chat-gin.com/router"
	"chat-gin.com/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	r := router.RouterInit()
	r.Run(":8081") // listen and serve on localhost:8081 (defult for windows "localhost:8080")
}
