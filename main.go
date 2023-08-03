package main

import (
	"fmt"

	"chat-gin.com/cmd"
	"chat-gin.com/router"
	"chat-gin.com/utils"
)

func main() {
	err:=cmd.RunSwagInit()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
	utils.InitConfig()
	utils.InitMySQL()
	r := router.RouterInit()
	r.Run(":8081") // listen and serve on localhost:8081 (defult for windows "localhost:8080")
}
