package main

import (
	"go_chat/router"
	"go_chat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySql()
	r := router.Router()
	
	r.Run(":8080")
}
