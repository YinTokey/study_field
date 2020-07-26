package routers

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("view/*")

	router.GET("/register",controllers.RegisterGet)

	return router;

}
