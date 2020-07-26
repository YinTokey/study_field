package main

import (
	"GitChat/database"
	"GitChat/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	database.InitMySql()

	router := routers.InitRouter()

	router.Static("/static","./static")

	router.Run(":8080")


}


func loginEndpoint(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest") //可设置默认值
	c.String(http.StatusOK, fmt.Sprintf("Hello %s \n", name))
}

func submitEndpoint(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest") //可设置默认值
	c.String(http.StatusOK, fmt.Sprintf("Hello %s \n", name))
}

func readEndpoint(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest") //可设置默认值
	c.String(http.StatusOK, fmt.Sprintf("Hello %s \n", name))
}
