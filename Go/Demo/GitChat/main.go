package main

import (
	"gitchat/database"
	"gitchat/routers"

)

func main() {

	database.InitMySql()

	router := routers.InitRouter()

	router.Static("/static","./static")

	router.Run(":8080")


}
