package router

import (
	"../controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("view/*")

	router.GET("/register",controller.RegisterGet)

	return router;

}
